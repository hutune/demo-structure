package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"mtsgn-access-user-svc/internal/models"
)

type UserStatus string

const (
	UserStatusActive     UserStatus = "active"
	UserStatusInactive   UserStatus = "inactive"
	UserStatusDisabled   UserStatus = "disabled"
	UserStatusUnverified UserStatus = "unverified"
)

type DateConfig struct {
	Format string `json:"format" gorm:"column:format"`
}

type TimeConfig struct {
	Format string `json:"format" gorm:"column:format"`
}

type User struct {
	models.BaseModel
	Email             string         `gorm:"type:varchar(255);column:email;not null;uniqueIndex:idx_email" json:"email"`
	FullName          string         `gorm:"type:varchar(255);column:full_name" json:"fullName"`
	FirstName         string         `gorm:"type:varchar(100);column:first_name" json:"firstName"`
	LastName          string         `gorm:"type:varchar(100);column:last_name" json:"lastName"`
	ProfileURL        string         `gorm:"type:text;column:profile_url" json:"profileUrl"`
	PhoneNumber       string         `gorm:"type:varchar(20);column:phone_number" json:"phoneNumber"`
	AddressLine       string         `gorm:"type:text;column:address_line" json:"addressLine"`
	Country           string         `gorm:"type:varchar(100);column:country" json:"country"`
	IsSuperAdmin      bool           `gorm:"type:boolean;column:is_super_admin;default:false" json:"isSuperAdmin"`
	IsVerified        bool           `gorm:"type:boolean;column:is_verified;default:false" json:"isVerified"`
	Gender            string         `gorm:"type:varchar(10);column:gender" json:"gender"`
	BirthDate         *time.Time     `gorm:"type:timestamp;column:birth_date" json:"birthDate"`
	Status            UserStatus     `gorm:"type:varchar(20);column:status;not null;default:'unverified'" json:"status"`
	IsCompanyInActive bool           `gorm:"type:boolean;column:is_company_in_active;default:false" json:"isCompanyInActive"`
	Metadata          datatypes.JSON `gorm:"type:jsonb;serializer:json" json:"metadata"`
}

func (User) TableName() string {
	return "users"
}

type UserConfig struct {
	models.BaseModel
	UserID       uuid.UUID      `gorm:"type:uuid;column:user_id;not null;index:idx_user_id" json:"userId"`
	LanguageCode string         `gorm:"type:varchar(2);column:language_code" json:"languageCode"`
	Date         DateConfig     `gorm:"type:jsonb;serializer:json;column:date_config" json:"date"`
	Time         TimeConfig     `gorm:"type:jsonb;serializer:json;column:time_config" json:"time"`
	Metadata     datatypes.JSON `gorm:"type:jsonb;serializer:json;column:metadata" json:"metadata"`
}

func (UserConfig) TableName() string {
	return "user_configs"
}
