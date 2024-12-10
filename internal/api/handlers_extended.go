package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/types"
)

// HandleGeneratePerson handles requests for generating person data
func (h *Handler) HandleGeneratePerson(c *gin.Context) {
	var req types.PersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GeneratePerson()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateNames handles requests for generating names
func (h *Handler) HandleGenerateNames(c *gin.Context) {
    var req types.CountRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    names := make([]string, req.Count)
    for i := 0; i < req.Count; i++ {
        names[i] = h.generator.GenerateName()
    }

    c.JSON(http.StatusOK, types.Response{
        Data: names,
        Meta: &types.MetaData{
            Count:     req.Count,
            TimeStamp: time.Now().Format(time.RFC3339),
        },
    })
}

// HandleGenerateDetailedProducts handles requests for generating detailed product data
func (h *Handler) HandleGenerateDetailedProducts(c *gin.Context) {
	var req types.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GenerateProductDetailed()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateVehicles handles requests for generating vehicle data
func (h *Handler) HandleGenerateVehicles(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GenerateVehicle()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGeneratePayments handles requests for generating payment data
func (h *Handler) HandleGeneratePayments(c *gin.Context) {
	var req types.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GeneratePaymentInfo()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateLocations handles requests for generating location data
func (h *Handler) HandleGenerateLocations(c *gin.Context) {
	var req types.LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GenerateLocation()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateFiles handles requests for generating file information
func (h *Handler) HandleGenerateFiles(c *gin.Context) {
	var req types.CountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GenerateFileInfo()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateInternetInfo handles requests for generating internet information
func (h *Handler) HandleGenerateInternetInfo(c *gin.Context) {
	var req types.InternetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := h.generator.GenerateMultiple(req.Count, func() interface{} {
		return h.generator.GenerateInternetDetailed()
	})

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}

// HandleGenerateCustom handles custom data generation requests
func (h *Handler) HandleGenerateCustom(c *gin.Context) {
	var req struct {
		Count      int                    `json:"count" binding:"required,gt=0"`
		Fields     map[string]string      `json:"fields" binding:"required"`
		Options    map[string]interface{} `json:"options,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	results := make([]map[string]interface{}, req.Count)

	for i := 0; i < req.Count; i++ {
		item := make(map[string]interface{})
		for field, fieldType := range req.Fields {
			switch fieldType {
			case "name":
				item[field] = h.generator.GenerateName()
			case "email":
				item[field] = h.generator.GenerateEmail()
			case "phone":
				item[field] = h.generator.GeneratePhone()
			case "address":
				item[field] = h.generator.GenerateAddress()
			case "company":
				item[field] = h.generator.GenerateCompany()
			case "job":
				item[field] = h.generator.GenerateJobDetails()
			case "product":
				item[field] = h.generator.GenerateProductDetailed()
			default:
				item[field] = h.generator.GenerateString(10) // default to random string
			}
		}
		results[i] = item
	}

	c.JSON(http.StatusOK, types.Response{
		Data: results,
		Meta: &types.MetaData{
			Count:      req.Count,
			TimeStamp:  time.Now().Format(time.RFC3339),
			Parameters: req,
			Generation: time.Since(start).String(),
		},
	})
}