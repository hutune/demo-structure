package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"mtsgn-access-user-svc/internal/models/user"
)

// Pagination represents pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

// Request DTOs
type GetUserProfileReq struct {
	UserID uuid.UUID `json:"userId" binding:"required,uuid"`
}

type UpdateProfileReq struct {
	UserID      uuid.UUID  `json:"userId" binding:"required,uuid"`
	FullName    *string    `json:"fullName"`
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	ProfileURL  *string    `json:"profileUrl"`
	PhoneNumber *string    `json:"phoneNumber"`
	AddressLine *string    `json:"addressLine"`
	Country     *string    `json:"country"`
	Gender      *string    `json:"gender"`
	BirthDate   *time.Time `json:"birthDate"`
}

type GetUsersReq struct {
	Page   int    `form:"page,default=1" binding:"min=1"`
	Limit  int    `form:"limit,default=10" binding:"min=1,max=100"`
	Search string `form:"search"`
	Status string `form:"status"`
}

type GetCompanyUsersReq struct {
	CompanyID uuid.UUID `json:"companyId" binding:"required,uuid"`
	Page      int       `form:"page,default=1" binding:"min=1"`
	Limit     int       `form:"limit,default=10" binding:"min=1,max=100"`
	Search    string    `form:"search"`
	Status    string    `form:"status"`
}

type GetStoreUsersReq struct {
	StoreID uuid.UUID `json:"storeId" binding:"required,uuid"`
	Page    int       `form:"page,default=1" binding:"min=1"`
	Limit   int       `form:"limit,default=10" binding:"min=1,max=100"`
	Search  string    `form:"search"`
	Status  string    `form:"status"`
}

type UpdateCompanyUserReq struct {
	UserID    uuid.UUID `json:"userId" binding:"required,uuid"`
	CompanyID uuid.UUID `json:"companyId" binding:"required,uuid"`
	Status    *string   `json:"status"`
}

type UpdateUserAdminReq struct {
	UserID uuid.UUID `json:"userId" binding:"required,uuid"`
	Status *string   `json:"status"`
}

// Response DTOs
type UserProfileResp struct {
	user.User
	Config *user.UserConfig `json:"config"`
}

type UserListItem struct {
	ID              uuid.UUID  `json:"id"`
	Email           string     `json:"email"`
	FullName        string     `json:"fullName"`
	FirstName       string     `json:"firstName"`
	LastName        string     `json:"lastName"`
	ProfileURL      string     `json:"profileUrl"`
	PhoneNumber     string     `json:"phoneNumber"`
	AddressLine     string     `json:"addressLine"`
	Country         string     `json:"country"`
	IsSuperAdmin    bool       `json:"isSuperAdmin"`
	IsVerified      bool       `json:"isVerified"`
	Gender          string     `json:"gender"`
	BirthDate       *time.Time `json:"birthDate"`
	Status          string     `json:"status"`
	IsCompanyActive bool       `json:"isCompanyInActive"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

type UsersResp struct {
	Users      []UserListItem `json:"users"`
	Pagination Pagination     `json:"pagination"`
}

type UserConfigResp struct {
	user.UserConfig
}

type UpdateUserResp struct {
	user.User
}

type UpdateStatusResp struct {
	Status string `json:"status"`
}

// User Config DTOs
type DateConfig struct {
	Format string `json:"format"`
}

type TimeConfig struct {
	Format string `json:"format"`
}

type UpdateUserConfigReq struct {
	UserID       uuid.UUID       `json:"userId" binding:"required,uuid"`
	LanguageCode *string         `json:"languageCode"`
	DateConfig   *DateConfig     `json:"date"`
	TimeConfig   *TimeConfig     `json:"time"`
	Metadata     *datatypes.JSON `json:"metadata"`
}

type GetUserConfigReq struct {
	UserID uuid.UUID `json:"userId" binding:"required,uuid"`
}
