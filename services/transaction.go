package services

import (
	"context"
	"errors"
	"one_pay/dto"
	"one_pay/model"
	"one_pay/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TransactionService interface {
	Create(ctx context.Context, req dto.TransactionDTO, userCode string) error
	GetByUserCode(ctx context.Context, userCode string) (*[]model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepo
	userRepo        repository.UserRepository
}

func NewTransactionService(transRepo repository.TransactionRepo, userRepo repository.UserRepository) TransactionService {
	return &transactionService{
		transactionRepo: transRepo,
		userRepo:        userRepo,
	}
}

func (t *transactionService) Create(ctx context.Context, req dto.TransactionDTO, userCode string) error {

	user, err := t.userRepo.FindByUserCode(ctx, req.UserCode)
	if err != nil {
		logrus.Fatalf("error : %s", err)
		return err
	}

	if user == nil {
		err = errors.New("bad request")
		logrus.Fatalf("error : %s", err)
		return err
	}

	// for extra security we will match the usercode in payload with token
	if userCode != user.UserCode {
		err = errors.New("unathorized user")
		logrus.Fatalf("error : %s", err)
		return err
	}

	trans := model.Transaction{
		UserCode:        req.UserCode,
		TransactionCode: uuid.NewString(),
		Amount:          req.Amount,
		PaymentMethod:   req.PaymentMethod,
		Currency:        req.Currency,
	}

	err = t.transactionRepo.CreateTransaction(ctx, &trans)
	if err != nil {
		logrus.Fatalf("error : %s", err)
		return err
	}

	return nil
}

func (t *transactionService) GetByUserCode(ctx context.Context, userCode string) (*[]model.Transaction, error) {

	user, err := t.userRepo.FindByUserCode(ctx, userCode)
	if err != nil {
		logrus.Fatalf("error : %s", err)
		return nil, err
	}

	if user == nil {
		err = errors.New("bad request")
		logrus.Fatalf("error : %s", err)
		return nil, err
	}

	data, err := t.transactionRepo.FindByUserCode(ctx, userCode)
	if err != nil {
		logrus.Fatalf("error : %s", err)
		return nil, err
	}

	return data, nil
}
