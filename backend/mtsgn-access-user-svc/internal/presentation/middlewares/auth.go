package middleware

import (
	"mtsgn-access-user-svc/internal/dto"

	"gitea.solu-m.io/smart-pos/sp-system-common-svc/common/constant"

	"gitea.solu-m.io/smart-pos/sp-system-common-svc/common/apperror"
	"github.com/gin-gonic/gin"
)

func AuthCompanyID() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader(constant.LocaleHeader)
		cid := c.GetHeader(constant.CompanyIDHeader)
		if cid == "" {
			dto.ErrorResponse(c, &apperror.ErrCompanyIDRequired, locale)
			c.Abort()
			return
		}
		c.Set(constant.CompanyIDHeader, cid)
		c.Next()
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader(constant.LocaleHeader)
		userID := c.GetHeader(constant.UserIDHeader)
		if userID == "" {
			dto.ErrorResponse(c, &apperror.ErrUserIDRequired, locale)
			c.Abort()
			return
		}
		c.Set(constant.UserIDHeader, userID)
		c.Next()
	}
}
