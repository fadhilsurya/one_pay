package controller

import (
	"context"
	"errors"
	"one_pay/dto"
	"one_pay/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) UserController {
	return &userController{
		userService: us,
	}

}

func (u *userController) Register(c *gin.Context) {

	var (
		rg   dto.RegisterDTO
		resp dto.ResponseDTO
	)

	if err := c.ShouldBind(&rg); err != nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("bad request").Error(),
			Success: false,
		}

		c.JSON(400, resp)
		return
	}

	err := u.userService.Register(c.Request.Context(), rg)
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
		Data:    nil,
	}
	c.JSON(200, resp)
}

func (u *userController) Login(c *gin.Context) {

	var (
		ur   dto.LoginDTO
		resp dto.ResponseDTO
	)

	ctx := context.Background()

	if err := c.ShouldBind(&ur); err != nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}

		c.JSON(400, resp)
		return
	}

	token, err := u.userService.Login(ctx, ur)
	if err != nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("internal server error").Error(),
			Success: false,
		}

		c.JSON(500, resp)
		return
	}

	if token == nil && err == nil {
		resp = dto.ResponseDTO{
			Data:    nil,
			Message: errors.New("phone number incorrect").Error(),
			Success: false,
		}

		c.JSON(400, resp)
		return
	}

	t := map[string]interface{}{
		"token": token,
	}

	resp = dto.ResponseDTO{
		Success: true,
		Message: "success",
		Data:    t,
	}

	c.JSON(200, resp)

}
