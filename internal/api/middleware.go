package api

import (
	"net/http"
	"sync"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/1lker/sd-gen-o2/internal/errors"
)

// RateLimiter implements a simple token bucket rate limiter
type RateLimiter struct {
	tokens    map[string][]time.Time
	rateLimit int
	mu        sync.Mutex
}

func NewRateLimiter(rateLimit int) *RateLimiter {
	return &RateLimiter{
		tokens:    make(map[string][]time.Time),
		rateLimit: rateLimit,
	}
}

// RateLimit middleware implements rate limiting per IP
func RateLimit(limiter *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now()
		windowStart := now.Add(-time.Minute)

		// Remove old tokens
		var validTokens []time.Time
		for _, t := range limiter.tokens[ip] {
			if t.After(windowStart) {
				validTokens = append(validTokens, t)
			}
		}

		if len(validTokens) >= limiter.rateLimit {
			c.JSON(http.StatusTooManyRequests, errors.RateLimitExceeded())
			c.Abort()
			return
		}

		limiter.tokens[ip] = append(validTokens, now)
		c.Next()
	}
}

// Logger middleware implements request logging
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ip := c.ClientIP()

		// Log in a structured format
		gin.DefaultWriter.Write([]byte(
			formatLog(method, path, query, status, latency, ip),
		))
	}
}

// Timeout middleware implements request timeout
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the request in a channel
		done := make(chan bool, 1)
		go func() {
			c.Next()
			done <- true
		}()

		// Wait for either completion or timeout
		select {
		case <-done:
			return
		case <-time.After(timeout):
			c.JSON(http.StatusRequestTimeout, errors.TimeoutError())
			c.Abort()
		}
	}
}

// CORS middleware implements Cross-Origin Resource Sharing
func CORS(allowedOrigins string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestSizeLimit middleware limits the request body size
func RequestSizeLimit(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}

// formatLog formats the log message
func formatLog(method, path, query string, status int, latency time.Duration, ip string) string {
	return fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s%s\n",
		time.Now().Format("2006/01/02 - 15:04:05"),
		status,
		latency,
		ip,
		method,
		path,
		func() string {
			if query != "" {
				return "?" + query
			}
			return ""
		}(),
	)
}