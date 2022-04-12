package main


import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GinGormRedisApi/api"
	"github.com/lisyaoran51/GinGormRedisApi/controller"
	"github.com/lisyaoran51/GinGormRedisApi/service"
	"github.com/lisyaoran51/GinGormRedisApi/repository"
	"github.com/lisyaoran51/GinGormRedisApi/counter"
	"github.com/lisyaoran51/GinGormRedisApi/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// var dsn = "jiayu:jiayu@tcp(127.0.0.1:3306)/go_test_models?charset=utf8mb4&parseTime=true"
// var dsn = "jiayu:jiayu@tcp(0.0.0.0:6033)/go_test_models?charset=utf8mb4&parseTime=true"
var dsn = "jiayu:jiayu@tcp(db:3306)/go_test_models?charset=utf8mb4&parseTime=true"

var (
	userRepository repository.UserRepository 	= repository.NewMysqlUserRepository(dsn)
	// callCounter counter.Counter 				= counter.NewRedisCounter("localhost:6379", 0)
	// callCounter counter.Counter 				= counter.NewRedisCounter("0.0.0.0:9736", 0)
	callCounter counter.Counter 				= counter.NewRedisCounter("redis:6379", 0)
	userService service.UserService 			= service.NewUserService(userRepository)
	userController controller.UserController 	= controller.NewUserController(userService)
)

// @in header
func main() {

	docs.SwaggerInfo.Title = "GinGormRedisAPI"
	docs.SwaggerInfo.Description = "GinGormRedisAPI - user insert api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "172.18.0.4:8000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	fmt.Println("start GinGormRedisApi...")

	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	userAPI := api.NewUserAPI(userController, callCounter)

	routes := server.Group(docs.SwaggerInfo.BasePath)
	{
		users := routes.Group("/users")
		{
			users.GET("", userAPI.GetUsers)
			// users.GET(":id", userAPI.GetUserByID)
			users.POST("", userAPI.CreateUser)
			users.PUT(":id", userAPI.UpdateUser)
			users.DELETE(":id", userAPI.DeleteUser)
		}

		count := routes.Group("/count")
		{
			count.GET("", userAPI.GetCount)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}