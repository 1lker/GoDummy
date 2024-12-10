package api

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/types"
)

// HandleGenerateTransactions handles requests for generating transaction data
func (h *Handler) HandleGenerateTransactions(c *gin.Context) {
	var req types.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: "Invalid request parameters: " + err.Error(),
		})
		return
	}

	// Validate amount range if provided
	if req.MaxAmount != 0 && req.MinAmount > req.MaxAmount {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: "minimum amount cannot be greater than maximum amount",
		})
		return
	}

	// Validate timestamp range if provided
	if req.MinTimestamp != "" && req.MaxTimestamp != "" {
		minTime, err1 := time.Parse(time.RFC3339, req.MinTimestamp)
		maxTime, err2 := time.Parse(time.RFC3339, req.MaxTimestamp)
		
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, types.Response{
				Error: "invalid timestamp format. Use RFC3339 format",
			})
			return
		}

		if minTime.After(maxTime) {
			c.JSON(http.StatusBadRequest, types.Response{
				Error: "minimum timestamp cannot be after maximum timestamp",
			})
			return
		}
	}

	// Generate transactions
	start := time.Now()
	transactions := make([]types.Transaction, req.Count)
	for i := 0; i < req.Count; i++ {
		transactions[i] = h.generator.GenerateTransaction(req)
	}

	// Prepare response
	c.JSON(http.StatusOK, types.Response{
		Data: transactions,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateBatchTransactions handles requests for generating batches of related transactions
func (h *Handler) HandleGenerateBatchTransactions(c *gin.Context) {
	var req types.BatchTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: "Invalid request parameters: " + err.Error(),
		})
		return
	}

	// Validate amount range if provided
	if req.MaxAmount != 0 && req.MinAmount > req.MaxAmount {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: "minimum amount cannot be greater than maximum amount",
		})
		return
	}

	start := time.Now()
	batches := make([][]types.Transaction, req.BatchCount)

	// Convert batch request to transaction request
	baseRequest := types.TransactionRequest{
		Count:     1, // Will be used for each transaction
		MinAmount: req.MinAmount,
		MaxAmount: req.MaxAmount,
		Currency:  req.Currency,
		Type:      req.Type,
		Status:    req.Status,
	}

	for i := 0; i < req.BatchCount; i++ {
		// Generate a batch of related transactions
		batch := make([]types.Transaction, req.BatchSize)
		
		// Generate parent transaction
		parentTx := h.generator.GenerateTransaction(baseRequest)
		batch[0] = parentTx

		// Generate child transactions
		for j := 1; j < req.BatchSize; j++ {
			// Link transactions through accounts
			childReq := baseRequest
			childReq.FromAccount = parentTx.To
			
			tx := h.generator.GenerateTransaction(childReq)
			tx.ParentTransactionID = parentTx.ID
			tx.Tags = append(tx.Tags, "child_transaction")
			batch[j] = tx
		}

		batches[i] = batch
	}

	c.JSON(http.StatusOK, types.Response{
		Data: batches,
		Meta: &types.MetaData{
			Count:      req.BatchCount * req.BatchSize,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}