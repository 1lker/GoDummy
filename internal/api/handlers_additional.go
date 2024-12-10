package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/types"
)

// HandleGenerateJob handles requests for generating job information
func (h *Handler) HandleGenerateJob(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateJobInfo()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}

// HandleGenerateProduct handles requests for generating product information
func (h *Handler) HandleGenerateProduct(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]interface{}, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateProduct()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}

// HandleGenerateCar handles requests for generating car information
func (h *Handler) HandleGenerateCar(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateCarInfo()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}

// HandleGenerateInternet handles requests for generating internet information
func (h *Handler) HandleGenerateInternet(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateInternetInfo()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}

// HandleGenerateFile handles requests for generating file information
func (h *Handler) HandleGenerateFile(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateFile()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}

// HandleGenerateColor handles requests for generating color information
func (h *Handler) HandleGenerateColor(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateColor()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	})
}