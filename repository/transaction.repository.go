package repository

import (
	"context"
	"one_pay/model"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, trans *model.Transaction) error
	FindByUserCode(ctx context.Context, userCode string) (*[]model.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (s *transactionRepo) CreateTransaction(ctx context.Context, trans *model.Transaction) error {

	return s.db.WithContext(ctx).Create(trans).Error
}

func (s *transactionRepo) FindByUserCode(ctx context.Context, userCode string) (*[]model.Transaction, error) {
	var (
		trans []model.Transaction
	)

	err := s.db.WithContext(ctx).Where("user_code = ?", userCode).Find(&trans).Error
	if err != nil {
		return nil, err
	}

	return &trans, nil
}
