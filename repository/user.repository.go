package repository

import (
	"context"
	"one_pay/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*[]model.User, error)
	FindByUserCode(ctx context.Context, userCode string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// Add your service methods here
func (s *userRepo) CreateUser(ctx context.Context, user *model.User) error {

	return s.db.WithContext(ctx).Create(user).Error
}

func (s *userRepo) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*[]model.User, error) {
	var (
		user []model.User
	)

	err := s.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userRepo) FindByUserCode(ctx context.Context, userCode string) (*model.User, error) {
	var (
		user model.User
	)

	err := s.db.WithContext(ctx).Where("user_code = ?", userCode).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
