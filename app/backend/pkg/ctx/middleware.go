package ctx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CompanyMiddleware is a middleware required company ID in the header
func CompanyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.GetHeader(CompanyIDHeader.Str())
		if cid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Company ID is required"})
			return
		}
		ctx := c.Request.Context()
		c.Set(CompanyIDHeader, cid)
		ctx = context.WithValue(ctx, CompanyIDHeader, cid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// UserMiddleware is a middleware required user ID in the header
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(UserIDHeader.Str())
		cid := c.GetHeader(CompanyIDHeader.Str())
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
			return
		}
		ctx := c.Request.Context()
		isAdmin := c.GetHeader(IsAdminHeader.Str())
		if isAdmin == "true" {
			c.Set(IsAdminHeader, true)
			ctx = context.WithValue(ctx, IsAdminHeader, true)
		} else {
			c.Set(IsAdminHeader, false)
			ctx = context.WithValue(ctx, IsAdminHeader, false)
		}

		ctx = context.WithValue(ctx, UserIDHeader, userID)
		if cid != "" {
			c.Set(CompanyIDHeader, cid)
			ctx = context.WithValue(ctx, CompanyIDHeader, cid)
		}
		c.Set(UserIDHeader, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.GetHeader(IsAdminHeader.Str())
		if isAdmin == "false" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Next()
	}
}
