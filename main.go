package main

import (
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.SetConfig("./config.yaml") // 临时
	if err := service.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/user/query", service.UserQuery)
	r.POST("/user/create", service.UserCreate)
	r.POST("/user/update", service.UserUpdate)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
