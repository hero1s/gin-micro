package routers

import (
	"gin-micro/api/controls"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户服务的所有接口
func RegisterUser(router *gin.Engine) {
	u := new(controls.UserControl)
	u.InitRoutes(router)

	user := router.Group("user")
	RegisterApiRoutes(user)
	RegisterAppRoutes(user)
	RegisterAuthRoutes(user)
}

// api路由注册
func RegisterApiRoutes(router *gin.RouterGroup) {
	api := router.Group("api")
	// 鉴权
	//api.Use(JWTAuth())
	api.POST("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello APP")
	})
}

// app 路由注册
func RegisterAppRoutes(router *gin.RouterGroup) {
	app := router.Group("app")
	// 鉴权
	//app.Use(auth.JWTAuth())
	app.GET("hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello APP")
	})
}

// 注册其他需要鉴权的接口
func RegisterAuthRoutes(router *gin.RouterGroup) {
	//router.Use(auth.JWTAuth())
	router.GET("any", func(context *gin.Context) {
		context.String(http.StatusOK, "say anything")
	})
}
