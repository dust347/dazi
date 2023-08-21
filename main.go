package main

import (
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := service.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/user/login", service.Login)
	r.Use(service.MiddlewareAuth())
	r.GET("/user/query", service.UserQuery)
	r.POST("/user/update", service.UserUpdate)
	r.POST("/nearby", service.Nearby)
	r.POST("/user/upload_avatar", service.UploadAvatar)

	r.Run()
}
