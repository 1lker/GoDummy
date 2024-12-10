package api

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/generator"
	"github.com/1lker/sd-gen-o2/internal/config"
)

type Server struct {
	router      *gin.Engine
	handler     *Handler
	config      *config.Config
	rateLimiter *RateLimiter
}

func NewServer(cfg *config.Config) *Server {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	
	// Create generator and handler
	gen := generator.New()
	handler := NewHandler(gen)
	
	// Create rate limiter
	rateLimiter := NewRateLimiter(cfg.API.RateLimit)

	return &Server{
		router:      router,
		handler:     handler,
		config:      cfg,
		rateLimiter: rateLimiter,
	}
}

func (s *Server) SetupMiddleware() {
	// Add recovery middleware
	s.router.Use(gin.Recovery())
	
	// Add custom logger
	s.router.Use(Logger())
	
	// Add rate limiting
	s.router.Use(RateLimit(s.rateLimiter))
	
	// Add CORS
	s.router.Use(CORS(s.config.Server.AllowedOrigins))
	
	// Add timeout
	timeout := time.Duration(s.config.API.TimeoutSeconds) * time.Second
	s.router.Use(Timeout(timeout))
	
	// Add request size limit
	s.router.Use(RequestSizeLimit(s.config.API.MaxRequestSize))
}

func (s *Server) SetupRoutes() {
	// Health check endpoint
	s.router.GET("/health", s.handler.HandleHealth)

	// API version group
	v1 := s.router.Group("/api/v1")
	{
		// Generator endpoints
		generate := v1.Group("/generate")
		{
			// Basic types
			generate.POST("/int", s.handler.HandleGenerateIntegers)
			generate.POST("/float", s.handler.HandleGenerateFloats)
			generate.POST("/string", s.handler.HandleGenerateStrings)
			generate.POST("/bool", s.handler.HandleGenerateBooleans)
			
			// Complex types
			generate.POST("/date", s.handler.HandleGenerateDates)
			generate.POST("/email", s.handler.HandleGenerateEmails)
			generate.POST("/phone", s.handler.HandleGeneratePhones)
			
			// New complex types
			generate.POST("/address", s.handler.HandleGenerateAddresses)
			generate.POST("/creditcard", s.handler.HandleGenerateCreditCards)
			generate.POST("/company", s.handler.HandleGenerateCompanies)
			
			// Batch generation
			generate.POST("/batch", s.handler.HandleBatchGenerate)
		}
	}
}

func (s *Server) Start() error {
	// Setup middleware and routes
	s.SetupMiddleware()
	s.SetupRoutes()

	// Start the server
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	fmt.Printf("Server starting on %s\n", addr)
	return s.router.Run(addr)
}