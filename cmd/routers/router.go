package routers

import (
	"test/cmd/handler"
	"test/cmd/middleware"
	"test/pkg/router"
)

func RegisterRouter(router *router.Router, userHandler handler.UserHandler) {
	router.Use(middleware.RequiredAPIKey)
	router.POST("/user/register", userHandler.Register)
	router.GET("/user/detail/:user_id", userHandler.Detail)
	router.GET("/user/list", userHandler.List)
}
