package dto

type TransactionDTO struct {
	UserCode      string  `json:"user_code" valid:"required"`
	Amount        float64 `json:"amount,string" valid:"required"`
	PaymentMethod string  `json:"payment_method" valid:"required"`
	PaymentStatus string  `json:"payment_status" valid:"required"`
	Currency      string  `json:"currency" valid:"required"`
}
