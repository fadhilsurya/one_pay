package controller

import (
	"context"
	"errors"
	"net/http"
	"one_pay/dto"
	"one_pay/helper"
	"one_pay/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransController interface {
	Create(c *gin.Context)
	GetTransactionHistory(c *gin.Context)
}

type transController struct {
	transService services.TransactionService
}

func NewTransController(ts services.TransactionService) TransController {
	return &transController{
		transService: ts,
	}

}

func (t *transController) Create(c *gin.Context) {

	var (
		tq   dto.TransactionDTO
		resp dto.ResponseDTO
	)

	ctx := context.Background()

	claims, exist := c.Get("token")
	if !exist {
		resp := dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("bad request").Error(),
			Success: false,
		}
		logrus.Fatal("Bad Request")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// decode token
	tokenClaims, ok := claims.(*helper.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
		return
	}

	if err := c.ShouldBind(&tq); err != nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("bad request").Error(),
			Success: false,
		}
		logrus.Fatal("Bad Requestn: %+v", err)
		c.JSON(400, resp)
		return
	}

	err := t.transService.Create(ctx, tq, tokenClaims.UserCode)
	if err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			resp = dto.ResponseDTO{
				Data:    nil,
				Message: errors.New("Record").Error(),
				Success: false,
			}

			c.JSON(400, resp)
			return
		}

		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("internal server error").Error(),
			Success: false,
		}

		c.JSON(500, resp)
		return
	}

	resp = dto.ResponseDTO{
		Success: true,
		Message: "success",
		Data:    nil,
	}
	c.JSON(200, resp)
}

func (t *transController) GetTransactionHistory(c *gin.Context) {

	var (
		tq   dto.TransactionDTO
		resp dto.ResponseDTO
	)

	ctx := context.Background()

	claims, exist := c.Get("token")
	if !exist {
		resp := dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("bad request").Error(),
			Success: false,
		}
		logrus.Fatal("Bad Request")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// decode token
	tokenClaims, ok := claims.(*helper.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
		return
	}

	if err := c.ShouldBind(&tq); err != nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("bad request").Error(),
			Success: false,
		}
		logrus.Fatal("Bad Request: %+v", err)
		c.JSON(400, resp)
		return
	}

	data, err := t.transService.GetByUserCode(ctx, tokenClaims.UserCode)
	if err != nil {

		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("internal server error").Error(),
			Success: false,
		}

		c.JSON(500, resp)
		return
	}

	resp = dto.ResponseDTO{
		Success: true,
		Message: "success",
		Data:    data,
	}
	c.JSON(200, resp)
}
