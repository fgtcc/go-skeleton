package routes

import (
	v1 "go-skeleton/api/v1"
	"go-skeleton/middleware"
	"go-skeleton/middleware/validator"

	"go-skeleton/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	v1.InitService()

	gin.SetMode(utils.AppMode)
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	router := r.Group("api/v1")
	{
		router.POST("login", validator.Validate(validator.LoginSchema()), v1.Login)
	}

	admin := r.Group("api/v1")
	admin.Use(middleware.AdminJwtToken())
	{
		admin.POST("user/add", validator.Validate(validator.AddUserSchema()), v1.AddUser)
		admin.POST("user/getList", validator.Validate(validator.GetUserListSchema()), v1.GetUserList)
		admin.POST("user/delete", validator.Validate(validator.DeleteUsersSchema()), v1.DeleteUsers)
		admin.POST("user/update", validator.Validate(validator.UpdateUserSchema()), v1.UpdateUser)
		admin.POST("user/batchUpdateStatus", validator.Validate(validator.BatchUpdateUserStatus()), v1.BatchUpdateUserStatus)
	}

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.POST("user/getInfo", validator.Validate(validator.GetUserInfoSchema()), v1.GetUserInfo)
	}

	return r
}
