package middleware

import (
	"os"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	// Security: Use environment variable to configure allowed origins
	// Set ALLOWED_ORIGINS environment variable with comma-separated URLs
	// Example: ALLOWED_ORIGINS=https://example.com,https://admin.example.com
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		origins := strings.Split(allowedOrigins, ",")
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
		config.AllowOrigins = origins
	} else {
		// Default: Allow all origins for backward compatibility
		// WARNING: This is insecure for production! Set ALLOWED_ORIGINS environment variable
		common.SysLog("WARNING: CORS AllowAllOrigins is enabled. Set ALLOWED_ORIGINS environment variable for production.")
		config.AllowAllOrigins = true
	}

	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// Security: Restrict allowed headers to only necessary ones
	config.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"New-Api-User",
		"X-Requested-With",
		"Accept",
		"Origin",
	}
	return cors.New(config)
}

func PoweredBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Security headers
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// HSTS - Only enable if using HTTPS
		if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		// CSP - Content Security Policy
		csp := "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline' 'unsafe-eval'; " +
			"style-src 'self' 'unsafe-inline'; " +
			"img-src 'self' data: https:; " +
			"font-src 'self' data:; " +
			"connect-src 'self' https:; " +
			"frame-ancestors 'none';"
		c.Header("Content-Security-Policy", csp)

		// Version header - only show in non-production or if DEBUG enabled
		if common.DebugEnabled || os.Getenv("GIN_MODE") != "release" {
			c.Header("X-New-Api-Version", common.Version)
		}

		c.Next()
	}
}
