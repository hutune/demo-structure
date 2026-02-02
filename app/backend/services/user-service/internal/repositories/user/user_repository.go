package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"rmn-backend/services/user-service/internal/common"
	"rmn-backend/services/user-service/internal/common/errors"
	"rmn-backend/services/user-service/internal/models/user"
)

type UserRepository struct {
	ctx *common.App
}

func NewUserRepository(ctx *common.App) *UserRepository {
	return &UserRepository{
		ctx: ctx,
	}
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var user user.User
	err := r.ctx.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

// GetPaginated retrieves users with pagination and filtering
func (r *UserRepository) GetPaginated(ctx context.Context, page, limit int, search, status string) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	query := r.ctx.DB.Model(&user.User{})

	// Apply filters
	if search != "" {
		query = query.Where("full_name ILIKE ? OR email ILIKE ?",
			fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetByCompanyID retrieves users for a specific company
func (r *UserRepository) GetByCompanyID(ctx context.Context, companyID uuid.UUID, page, limit int, search, status string) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	query := r.ctx.DB.Model(&user.User{}).Where("company_id = ?", companyID)

	// Apply additional filters
	if search != "" {
		query = query.Where("full_name ILIKE ? OR email ILIKE ?",
			fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetByStoreID retrieves users for a specific store
func (r *UserRepository) GetByStoreID(ctx context.Context, storeID uuid.UUID, page, limit int, search, status string) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	query := r.ctx.DB.Model(&user.User{}).Where("store_id = ?", storeID)

	// Apply additional filters
	if search != "" {
		query = query.Where("full_name ILIKE ? OR email ILIKE ?",
			fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	return r.ctx.DB.Create(user).Error
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	return r.ctx.DB.Save(user).Error
}

// GetConfigByUserID retrieves user configuration for a user
func (r *UserRepository) GetConfigByUserID(ctx context.Context, userID uuid.UUID) (*user.UserConfig, error) {
	var config user.UserConfig
	err := r.ctx.DB.Where("user_id = ?", userID).First(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &config, nil
}

// UpdateConfig updates user configuration
func (r *UserRepository) UpdateConfig(ctx context.Context, config *user.UserConfig) error {
	return r.ctx.DB.Save(config).Error
}
