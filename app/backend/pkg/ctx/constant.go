package ctx

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey string

const (
	CompanyIDHeader ctxKey = "X-Company-ID"
	UserIDHeader    ctxKey = "X-User-ID"
	IsAdminHeader   ctxKey = "X-Is-Admin"
	LocaleHeader    ctxKey = "Locale"
)

func (k ctxKey) Str() string {
	return string(k)
}

// Helper to retrieve UserID safely
func UserIDCtx(ctx context.Context) uuid.UUID {
	id := ctx.Value(UserIDHeader)
	if id == nil || id == "" {
		return uuid.Nil
	}
	return uuid.MustParse(id.(string))
}

func IsAdminCtx(ctx context.Context) bool {
	isAdmin, ok := ctx.Value(IsAdminHeader).(bool)
	if !ok {
		return false
	}
	return isAdmin
}

func CompanyIDCtx(ctx context.Context) uuid.UUID {
	companyID := ctx.Value(CompanyIDHeader)
	if companyID == nil || companyID == "" {
		return uuid.Nil
	}
	return uuid.MustParse(companyID.(string))
}

func UserIDGinCtx(ginCtx *gin.Context) uuid.UUID {
	userID, ok := ginCtx.Get(UserIDHeader)
	if !ok || userID == "" {
		return uuid.Nil
	}
	return uuid.MustParse(userID.(string))
}

func IsAdminGinCtx(ginCtx *gin.Context) bool {
	isAdmin, ok := ginCtx.Get(IsAdminHeader)
	if !ok || isAdmin == "" {
		return false
	}
	return isAdmin.(bool)
}

func CompanyIDGinCtx(ginCtx *gin.Context) uuid.UUID {
	companyID, ok := ginCtx.Get(CompanyIDHeader)
	if !ok || companyID == "" {
		return uuid.Nil
	}
	return uuid.MustParse(companyID.(string))
}

func MakeAuthHeader(ctx context.Context) map[string]string {
	userId := UserIDCtx(ctx)
	companyId := CompanyIDCtx(ctx)
	isAdmin := IsAdminCtx(ctx)
	headers := make(map[string]string)

	if userId != uuid.Nil {
		headers[UserIDHeader.Str()] = userId.String()
	}
	if companyId != uuid.Nil {
		headers[CompanyIDHeader.Str()] = companyId.String()
	}
	if isAdmin {
		headers[IsAdminHeader.Str()] = "true"
	}
	return headers
}
