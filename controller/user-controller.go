package controller

import (
	// "net/http"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GinGormRedisApi/service"
	"github.com/lisyaoran51/GinGormRedisApi/entity"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type userController struct{
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	c.service.Save(user)

	return nil
}

func (c *userController) Update(ctx *gin.Context) error {
	fmt.Println("--- userController.Update():")

	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	user.UserId = id

	err = c.service.Update(user)
	return err
}

func (c *userController) Delete(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.UserId = id

	err = c.service.Delete(user)
	return err
}