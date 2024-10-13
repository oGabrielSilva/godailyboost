package exceptions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Exception struct {
	Status    int       `json:"status"`
	Method    string    `json:"method"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	URL       string    `json:"URL"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *Exception) Error() string {
	return fmt.Sprintf("[%v] (%d) %s - %s -> %s", e.Timestamp, e.Status, e.Type, e.URL, e.Message)
}

// BadRequest (400)
func BadRequest(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusBadRequest,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "BAD_REQUEST",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// Unauthorized (401)
func Unauthorized(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusUnauthorized,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "UNAUTHORIZED",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// Forbidden (403)
func Forbidden(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusForbidden,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "FORBIDDEN",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// NotFound (404)
func NotFound(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusNotFound,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "NOT_FOUND",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// Conflict (409)
func Conflict(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusConflict,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "CONFLICT",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// InternalServerError (500)
func InternalServerError(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusInternalServerError,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "INTERNAL_SERVER_ERROR",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}

// ServiceUnavailable (503)
func ServiceUnavailable(message string, c *gin.Context) *Exception {
	return &Exception{
		Status:    http.StatusServiceUnavailable,
		Method:    c.Request.Method,
		Message:   message,
		Type:      "SERVICE_UNAVAILABLE",
		URL:       c.FullPath(),
		Timestamp: time.Now().UTC(),
	}
}
