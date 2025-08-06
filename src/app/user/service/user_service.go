package service

import (
	"errors"

	"github.com/faisd405/go-restapi-gin/src/app/user/model"
	"github.com/faisd405/go-restapi-gin/src/app/user/repository"
	"github.com/faisd405/go-restapi-gin/src/utils"
	"gorm.io/gorm"
)

type UserService interface {
	Register(req model.RegisterRequest) (*model.User, error)
	Login(req model.LoginRequest) (*model.LoginResponse, error)
	GetProfile(userID uint) (*model.UserResponse, error)
	UpdateProfile(userID uint, req model.UpdateUserRequest) (*model.UserResponse, error)
	ChangePassword(userID uint, req model.ChangePasswordRequest) error
	GetAllUsers(page, limit int) ([]model.UserResponse, int64, error)
	DeleteUser(userID uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(req model.RegisterRequest) (*model.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists with this email")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
		IsActive: true,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	}, nil
}

func (s *userService) GetProfile(userID uint) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

func (s *userService) UpdateProfile(userID uint, req model.UpdateUserRequest) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

func (s *userService) ChangePassword(userID uint, req model.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Check current password
	if !utils.CheckPassword(req.CurrentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// Update password
	return s.userRepo.UpdatePassword(userID, hashedPassword)
}

func (s *userService) GetAllUsers(page, limit int) ([]model.UserResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	users, total, err := s.userRepo.GetAll(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	userResponses := make([]model.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	return userResponses, total, nil
}

func (s *userService) DeleteUser(userID uint) error {
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(userID)
}
