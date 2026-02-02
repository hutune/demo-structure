package user

import (
	"context"
	"time"

	"github.com/google/uuid"

	"mtsgn-access-user-svc/internal/common/errors"
	userdto "mtsgn-access-user-svc/internal/dto/user"
	"mtsgn-access-user-svc/internal/models/user"
	userrepo "mtsgn-access-user-svc/internal/repositories/user"
)

type UserService struct {
	userRepo userrepo.UserRepository
}

func NewUserService(userRepo userrepo.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetUserProfile retrieves user profile by ID
func (s *UserService) GetUserProfile(ctx context.Context, req *userdto.GetUserProfileReq) (*userdto.UserProfileResp, error) {
	userModel, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	// Get user config
	userConfig, err := s.userRepo.GetConfigByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	resp := &userdto.UserProfileResp{
		User:   *userModel,
		Config: userConfig,
	}

	return resp, nil
}

// UpdateUserProfile updates user profile
func (s *UserService) UpdateUserProfile(ctx context.Context, req *userdto.UpdateProfileReq) (*userdto.UpdateUserResp, error) {
	// Validate user exists
	existingUser, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, errors.ErrNotFound
	}

	// Update fields
	if req.FullName != nil {
		existingUser.FullName = *req.FullName
	}
	if req.FirstName != nil {
		existingUser.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		existingUser.LastName = *req.LastName
	}
	if req.ProfileURL != nil {
		existingUser.ProfileURL = *req.ProfileURL
	}
	if req.PhoneNumber != nil {
		existingUser.PhoneNumber = *req.PhoneNumber
	}
	if req.AddressLine != nil {
		existingUser.AddressLine = *req.AddressLine
	}
	if req.Country != nil {
		existingUser.Country = *req.Country
	}
	if req.Gender != nil {
		existingUser.Gender = *req.Gender
	}
	if req.BirthDate != nil {
		existingUser.BirthDate = req.BirthDate
	}

	err = s.userRepo.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return &userdto.UpdateUserResp{User: *existingUser}, nil
}

// GetUsers retrieves users with pagination and filtering
func (s *UserService) GetUsers(ctx context.Context, req *userdto.GetUsersReq) (*userdto.UsersResp, error) {
	users, total, err := s.userRepo.GetPaginated(ctx, req.Page, req.Limit, req.Search, req.Status)
	if err != nil {
		return nil, err
	}

	userList := make([]userdto.UserListItem, len(users))
	for i, user := range users {
		var createdAt, updatedAt time.Time
		if user.CreatedAt != nil {
			createdAt = *user.CreatedAt
		}
		if user.UpdatedAt != nil {
			updatedAt = *user.UpdatedAt
		}
		userList[i] = userdto.UserListItem{
			ID:              user.ID,
			Email:           user.Email,
			FullName:        user.FullName,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			ProfileURL:      user.ProfileURL,
			PhoneNumber:     user.PhoneNumber,
			AddressLine:     user.AddressLine,
			Country:         user.Country,
			IsSuperAdmin:    user.IsSuperAdmin,
			IsVerified:      user.IsVerified,
			Gender:          user.Gender,
			BirthDate:       user.BirthDate,
			Status:          string(user.Status),
			IsCompanyActive: user.IsCompanyInActive,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
	}

	totalPages := int(total) / req.Limit
	if int(total)%req.Limit > 0 {
		totalPages++
	}

	return &userdto.UsersResp{
		Users: userList,
		Pagination: userdto.Pagination{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// GetCompanyUsers retrieves users for a specific company
func (s *UserService) GetCompanyUsers(ctx context.Context, req *userdto.GetCompanyUsersReq) (*userdto.UsersResp, error) {
	users, total, err := s.userRepo.GetByCompanyID(ctx, req.CompanyID, req.Page, req.Limit, req.Search, req.Status)
	if err != nil {
		return nil, err
	}

	userList := make([]userdto.UserListItem, len(users))
	for i, user := range users {
		var createdAt, updatedAt time.Time
		if user.CreatedAt != nil {
			createdAt = *user.CreatedAt
		}
		if user.UpdatedAt != nil {
			updatedAt = *user.UpdatedAt
		}
		userList[i] = userdto.UserListItem{
			ID:              user.ID,
			Email:           user.Email,
			FullName:        user.FullName,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			ProfileURL:      user.ProfileURL,
			PhoneNumber:     user.PhoneNumber,
			AddressLine:     user.AddressLine,
			Country:         user.Country,
			IsSuperAdmin:    user.IsSuperAdmin,
			IsVerified:      user.IsVerified,
			Gender:          user.Gender,
			BirthDate:       user.BirthDate,
			Status:          string(user.Status),
			IsCompanyActive: user.IsCompanyInActive,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
	}

	totalPages := int(total) / req.Limit
	if int(total)%req.Limit > 0 {
		totalPages++
	}

	return &userdto.UsersResp{
		Users: userList,
		Pagination: userdto.Pagination{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// GetStoreUsers retrieves users for a specific store
func (s *UserService) GetStoreUsers(ctx context.Context, req *userdto.GetStoreUsersReq) (*userdto.UsersResp, error) {
	users, total, err := s.userRepo.GetByStoreID(ctx, req.StoreID, req.Page, req.Limit, req.Search, req.Status)
	if err != nil {
		return nil, err
	}

	userList := make([]userdto.UserListItem, len(users))
	for i, user := range users {
		var createdAt, updatedAt time.Time
		if user.CreatedAt != nil {
			createdAt = *user.CreatedAt
		}
		if user.UpdatedAt != nil {
			updatedAt = *user.UpdatedAt
		}
		userList[i] = userdto.UserListItem{
			ID:              user.ID,
			Email:           user.Email,
			FullName:        user.FullName,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			ProfileURL:      user.ProfileURL,
			PhoneNumber:     user.PhoneNumber,
			AddressLine:     user.AddressLine,
			Country:         user.Country,
			IsSuperAdmin:    user.IsSuperAdmin,
			IsVerified:      user.IsVerified,
			Gender:          user.Gender,
			BirthDate:       user.BirthDate,
			Status:          string(user.Status),
			IsCompanyActive: user.IsCompanyInActive,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}
	}

	totalPages := int(total) / req.Limit
	if int(total)%req.Limit > 0 {
		totalPages++
	}

	return &userdto.UsersResp{
		Users: userList,
		Pagination: userdto.Pagination{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// UpdateCompanyUser updates user status for a company
func (s *UserService) UpdateCompanyUser(ctx context.Context, req *userdto.UpdateCompanyUserReq) error {
	userModel, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if userModel == nil {
		return errors.ErrNotFound
	}

	// Update company association and status
	if req.Status != nil {
		userModel.Status = user.UserStatus(*req.Status)
	}

	// TODO: Implement company association logic
	// This would involve linking user to company through a junction table

	return s.userRepo.Update(ctx, userModel)
}

// UpdateUserAdmin updates user status (admin only)
func (s *UserService) UpdateUserAdmin(ctx context.Context, req *userdto.UpdateUserAdminReq) error {
	userModel, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if userModel == nil {
		return errors.ErrNotFound
	}

	// Update status
	if req.Status != nil {
		userModel.Status = user.UserStatus(*req.Status)
	}

	return s.userRepo.Update(ctx, userModel)
}

// GetUserConfig retrieves user configuration
func (s *UserService) GetUserConfig(ctx context.Context, req *userdto.GetUserConfigReq) (*userdto.UserConfigResp, error) {
	userConfigModel, err := s.userRepo.GetConfigByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &userdto.UserConfigResp{UserConfig: *userConfigModel}, nil
}

// UpdateUserConfig updates user configuration
func (s *UserService) UpdateUserConfig(ctx context.Context, req *userdto.UpdateUserConfigReq) error {
	userConfigModel, err := s.userRepo.GetConfigByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if userConfigModel == nil {
		// Create new config
		userConfigModel = &user.UserConfig{
			UserID: req.UserID,
		}
	}

	// Update fields
	if req.LanguageCode != nil {
		userConfigModel.LanguageCode = *req.LanguageCode
	}
	if req.DateConfig != nil {
		userConfigModel.Date = user.DateConfig{
			Format: req.DateConfig.Format,
		}
	}
	if req.TimeConfig != nil {
		userConfigModel.Time = user.TimeConfig{
			Format: req.TimeConfig.Format,
		}
	}
	if req.Metadata != nil {
		userConfigModel.Metadata = *req.Metadata
	}

	return s.userRepo.UpdateConfig(ctx, userConfigModel)
}

// HandleUserCreated handles user creation events from auth service
func (s *UserService) HandleUserCreated(ctx context.Context, userID uuid.UUID, email string) error {
	// TODO: Implement user profile initialization
	// This would set default values for user profile based on system settings
	return nil
}
