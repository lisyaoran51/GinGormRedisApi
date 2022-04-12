package api

import (
	"net/http"
	"fmt"
	// "encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GinGormRedisApi/controller"
	"github.com/lisyaoran51/GinGormRedisApi/dto"
	"github.com/lisyaoran51/GinGormRedisApi/counter"


)


type UserApi struct {
	userController controller.UserController
	callCounter counter.Counter
}

func NewUserAPI(userController controller.UserController, callCounter counter.Counter) *UserApi {
	return &UserApi {
		userController: userController,
		callCounter: callCounter,
	}
}

// GetUsers godoc
// @Summary List existing users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} entity.User
// @Failure 401 {object} dto.Response
// @Router /users [get]
func (api *UserApi) GetUsers(ctx *gin.Context) {

	api.callCounter.Count()

	ctx.JSON(200, api.userController.FindAll())
}

// CreateUser godoc
// @Summary Create new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "Create user"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /users [post]
func (api *UserApi) CreateUser(ctx *gin.Context) {

	api.callCounter.Count()

	err := api.userController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "created",
		})
	}
}

// UpdateUser godoc
// @Summary Update users
// @Description Update a single user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "user ID"
// @Param user body entity.User true "Update user"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /users/{id} [put]
func (api *UserApi) UpdateUser(ctx *gin.Context) {

	api.callCounter.Count()

	err := api.userController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "updated",
		})
	}
}

// DeleteUser godoc
// @Summary Remove users
// @Description Delete a single user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "user ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /users/{id} [delete]
func (api *UserApi) DeleteUser(ctx *gin.Context) {

	api.callCounter.Count()

	err := api.userController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "deleted",
		})
	}
}


// GetCount godoc
// @Summary count all api calls
// @Description count all api calls
// @Tags count
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int
// @Router /count [post]
func (api *UserApi) GetCount(ctx *gin.Context) {

	count := api.callCounter.GetCount()
	
	m := map[string]int{
        "count": count,
    }
	// jsonString, _ := json.Marshal(m)

	fmt.Println("UserApi.GetCount(): temp api call count", count)

	ctx.JSON(http.StatusOK, m)
	// ctx.JSON(http.StatusOK, &dto.Response{
	// 	Message: string(jsonString),
	// })
}