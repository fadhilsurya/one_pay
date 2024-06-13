package services

import (
	"context"
	"errors"
	"one_pay/dto"
	"one_pay/helper"
	"one_pay/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*[]model.User, error) {
	args := m.Called(ctx, phoneNumber)
	if args.Get(0) != nil {
		return args.Get(0).(*[]model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByUserCode(ctx context.Context, userCode string) (*model.User, error) {
	args := m.Called(ctx, userCode)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestUserService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("success", func(t *testing.T) {
		req := dto.RegisterDTO{
			FullName:    "John Doe",
			Username:    "johndoe",
			Address:     "123 Main St",
			PhoneNumber: "1234567890",
			Role:        "user",
			Password:    "password",
		}

		userMatcher := mock.MatchedBy(func(user *model.User) bool {
			return user.FullName == req.FullName &&
				user.Username == req.Username &&
				user.Address == req.Address &&
				user.PhoneNumber == req.PhoneNumber &&
				user.Role == req.Role &&
				helper.CheckPasswordHash(req.Password, user.Password)
		})

		mockRepo.On("CreateUser", mock.Anything, userMatcher).Return(nil)

		err := service.Register(context.Background(), req)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		req := dto.RegisterDTO{
			FullName:    "John Doe",
			Username:    "johndoe",
			Address:     "123 Main St",
			PhoneNumber: "1234567890",
			Role:        "user",
		}

		userMatcher := mock.MatchedBy(func(user *model.User) bool {
			return user.FullName == req.FullName &&
				user.Username == req.Username &&
				user.Address == req.Address &&
				user.PhoneNumber == req.PhoneNumber &&
				user.Role == req.Role &&
				helper.CheckPasswordHash(req.Password, user.Password)
		})

		mockRepo.On("CreateUser", mock.Anything, userMatcher).Return(errors.New("failed to create user"))

		err := service.Register(context.Background(), req)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to create user", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	hashedPassword, _ := helper.HashPassword("password")

	t.Run("success", func(t *testing.T) {
		req := dto.LoginDTO{
			PhoneNumber: "1234567890",
			Password:    "password",
		}

		user := model.User{
			FullName:    "John Doe",
			Username:    "johndoe",
			Address:     "123 Main St",
			PhoneNumber: "1234567890",
			Role:        "user",
			UserCode:    uuid.NewString(),
			Password:    hashedPassword,
		}

		mockRepo.On("FindByPhoneNumber", mock.Anything, req.PhoneNumber).Return(&[]model.User{user}, nil)

		token, err := service.Login(context.Background(), req)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		req := dto.LoginDTO{
			PhoneNumber: "1234567890",
			Password:    "password",
		}

		mockRepo.On("FindByPhoneNumber", mock.Anything, req.PhoneNumber).Return(nil, nil)

		_, err := service.Login(context.Background(), req)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

}
