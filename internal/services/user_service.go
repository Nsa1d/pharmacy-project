package services

import (
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"strings"
)

type UserService interface {
	CreateUser(req models.UserCreateRequest) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	users repository.UserRepository
}

func NewUserService(
	users repository.UserRepository,
) UserService {
	return &userService{
		users: users,
	}
}

func (u *userService) CreateUser(req models.UserCreateRequest) (*models.User, error) {

	if err := u.validateUserCreate(req); err != nil {
		return nil, err
	}

	fullName := strings.TrimSpace(req.FullName)
	email := strings.ToLower(strings.TrimSpace(req.Email))
	phone := strings.TrimSpace(req.Phone)
	address := strings.TrimSpace(req.DefaultAddress)

	existingEmail, err := u.users.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if existingEmail != nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	existingPhone, err := u.users.GetByPhone(phone)
	if err != nil {
		return nil, err
	}

	if existingPhone != nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	user := &models.User{
		FullName:       fullName,
		Email:          email,
		Phone:          phone,
		DefaultAddress: address,
	}

	if err := u.users.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := u.users.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}
	return user, nil
}

func (u *userService) validateUserCreate(req models.UserCreateRequest) error {
	fullName := strings.TrimSpace(req.FullName)
	email := strings.TrimSpace(req.Email)
	phone := strings.TrimSpace(req.Phone)
	address := strings.TrimSpace(req.DefaultAddress)

	if fullName == "" {
		return apperrors.ErrInvalidUserInput
	}
	if len([]rune(fullName)) < 2 {
		return apperrors.ErrInvalidUserInput
	}

	if email == "" {
		return apperrors.ErrInvalidUserInput
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return apperrors.ErrInvalidUserInput
	}

	if phone == "" {
		return apperrors.ErrInvalidUserInput
	}
	if len([]rune(phone)) < 10 || len([]rune(phone)) > 20 {
		return apperrors.ErrInvalidUserInput
	}

	if address == "" {
		return apperrors.ErrInvalidUserInput
	}
	if len([]rune(address)) < 5 {
		return apperrors.ErrInvalidUserInput
	}

	return nil
}
