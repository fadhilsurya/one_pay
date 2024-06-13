package services

import (
	"context"
	"errors"
	"one_pay/dto"
	"one_pay/helper"
	"one_pay/model"
	"one_pay/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	Register(ctx context.Context, req dto.RegisterDTO) error
	Login(ctx context.Context, req dto.LoginDTO) (interface{}, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, req dto.RegisterDTO) error {

	pass, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := model.User{
		FullName:    req.FullName,
		Username:    req.Username,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
		UserCode:    uuid.NewString(),
		Password:    pass,
	}

	err = s.userRepo.CreateUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req dto.LoginDTO) (interface{}, error) {

	userData, err := s.userRepo.FindByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		logrus.Fatalf("error : %+v", err)
		return nil, errors.New("server busy")
	}

	if userData == nil || len(*userData) == 0 {
		return nil, nil
	}

	u := (*userData)

	if !helper.CheckPasswordHash(req.Password, u[0].Password) {
		logrus.Fatalf("error : %+v", err)
		return nil, errors.New("incorrect password")
	}

	tokenJWT, err := helper.NewJWTHelper().GenerateJWT(u[0].UserCode)
	if err != nil {
		logrus.Fatalf("error : %+v", err)
		return nil, errors.New("server busy")
	}

	return tokenJWT, nil
}
