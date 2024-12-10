package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/generator"
	"github.com/1lker/sd-gen-o2/internal/types"
	"github.com/1lker/sd-gen-o2/internal/errors"
)

type Handler struct {
	generator *generator.Generator
}

func NewHandler(g *generator.Generator) *Handler {
	return &Handler{
		generator: g,
	}
}

// HandleGenerateIntegers handles requests for generating integers
func (h *Handler) HandleGenerateIntegers(c *gin.Context) {
	var req types.IntegerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: err.Error(),
		})
		return
	}

	results := make([]int, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateInt(req.Min, req.Max)
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
	})
}

// HandleGenerateStrings handles requests for generating strings
func (h *Handler) HandleGenerateStrings(c *gin.Context) {
	var req types.StringRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: err.Error(),
		})
		return
	}

	results := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateString(req.Length)
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
	})
}

// HandleGenerateBooleans handles requests for generating booleans
func (h *Handler) HandleGenerateBooleans(c *gin.Context) {
	var req types.BooleanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Error: err.Error(),
		})
		return
	}

	results := make([]bool, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateBool()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
	})
}

// HandleGenerateFloats handles requests for generating floats
func (h *Handler) HandleGenerateFloats(c *gin.Context) {
	var req types.FloatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	results := make([]float64, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = h.generator.GenerateFloat(req.Min, req.Max)
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
	})
}

// HandleGenerateDates handles requests for generating dates
func (h *Handler) HandleGenerateDates(c *gin.Context) {
	var req types.DateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	dates := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		dates[i] = h.generator.GenerateDate(req.StartYear, req.EndYear)
	}

	c.JSON(http.StatusOK, types.Response{
		Data: dates,
	})
}

// HandleGenerateEmails handles requests for generating email addresses
func (h *Handler) HandleGenerateEmails(c *gin.Context) {
	var req types.EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	emails := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		emails[i] = h.generator.GenerateEmail()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: emails,
	})
}

// HandleGeneratePhones handles requests for generating phone numbers
func (h *Handler) HandleGeneratePhones(c *gin.Context) {
	var req types.EmailRequest // We can reuse the Count-only request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	phones := make([]string, req.Count)
	for i := 0; i < req.Count; i++ {
		phones[i] = h.generator.GeneratePhone()
	}

	c.JSON(http.StatusOK, types.Response{
		Data: phones,
	})
}

// HandleHealth handles health check requests
func (h *Handler) HandleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"time":   time.Now().Format(time.RFC3339),
	})
}



// HandleGenerateAddresses handles requests for generating addresses
func (h *Handler) HandleGenerateAddresses(c *gin.Context) {
	var req types.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	addresses := make([]types.Address, req.Count)
	for i := 0; i < req.Count; i++ {
		addresses[i] = h.generator.GenerateAddress()
		
		// Apply filters if specified
		if req.Country != "" && addresses[i].Country != req.Country {
			i-- // Regenerate if doesn't match filter
			continue
		}
		if req.State != "" && addresses[i].State != req.State {
			i-- // Regenerate if doesn't match filter
			continue
		}
	}

	c.JSON(http.StatusOK, types.Response{
		Data: addresses,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
			Parameters: req,
		},
	})
}

// HandleGenerateCreditCards handles requests for generating credit cards
func (h *Handler) HandleGenerateCreditCards(c *gin.Context) {
	var req types.CreditCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	cards := make([]types.CreditCard, req.Count)
	for i := 0; i < req.Count; i++ {
		cards[i] = h.generator.GenerateCreditCard()
		
		// Apply filter if specified
		if req.CardType != "" && cards[i].Type != req.CardType {
			i-- // Regenerate if doesn't match filter
			continue
		}
	}

	c.JSON(http.StatusOK, types.Response{
		Data: cards,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
			Parameters: req,
		},
	})
}

// HandleGenerateCompanies handles requests for generating companies
func (h *Handler) HandleGenerateCompanies(c *gin.Context) {
	var req types.CompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	companies := make([]types.Company, req.Count)
	for i := 0; i < req.Count; i++ {
		companies[i] = h.generator.GenerateCompany()
		
		// Apply filter if specified
		if req.Industry != "" && companies[i].Industry != req.Industry {
			i-- // Regenerate if doesn't match filter
			continue
		}
	}

	c.JSON(http.StatusOK, types.Response{
		Data: companies,
		Meta: &types.MetaData{
			Count:     req.Count,
			TimeStamp: time.Now().Format(time.RFC3339),
			Parameters: req,
		},
	})
}

// HandleBatchGenerate handles batch generation requests
func (h *Handler) HandleBatchGenerate(c *gin.Context) {
	var req types.BatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidInput("Invalid request parameters", err.Error()))
		return
	}

	results := make(map[string]interface{})
	for _, genReq := range req.Requests {
		var result interface{}
		var err error

		switch genReq.Type {
		case "address":
			options := types.AddressRequest{Count: 1}
			if genReq.Options != nil {
				// Parse options
			}
			addresses := make([]types.Address, options.Count)
			for i := 0; i < options.Count; i++ {
				addresses[i] = h.generator.GenerateAddress()
			}
			result = addresses

		case "creditcard":
			options := types.CreditCardRequest{Count: 1}
			if genReq.Options != nil {
				// Parse options
			}
			cards := make([]types.CreditCard, options.Count)
			for i := 0; i < options.Count; i++ {
				cards[i] = h.generator.GenerateCreditCard()
			}
			result = cards

		case "company":
			options := types.CompanyRequest{Count: 1}
			if genReq.Options != nil {
				// Parse options
			}
			companies := make([]types.Company, options.Count)
			for i := 0; i < options.Count; i++ {
				companies[i] = h.generator.GenerateCompany()
			}
			result = companies

		default:
			err = errors.InvalidInput("Unsupported type", genReq.Type)
		}

		if err != nil {
			results[genReq.Type] = map[string]string{"error": err.Error()}
		} else {
			results[genReq.Type] = result
		}
	}

c.JSON(http.StatusOK, types.Response{
	Data: results,
	Meta: &types.MetaData{
		Count:     len(req.Requests),
		TimeStamp: time.Now().Format(time.RFC3339),
		Types:     getRequestTypes(req.Requests),
	},
})
}

// getRequestTypes extracts unique types from batch requests
func getRequestTypes(requests []types.GenerationRequest) []string {
	typeMap := make(map[string]bool)
	for _, req := range requests {
		typeMap[req.Type] = true
	}

	types := make([]string, 0, len(typeMap))
	for t := range typeMap {
		types = append(types, t)
	}
	return types
}