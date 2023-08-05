package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasin-yumrutas/controller"
)

func main() {
	r := gin.Default()

	r.GET("/google/login", controller.GoogleLogin)
	r.GET("/google/callback", controller.GoogleCallBack)

	r.Run(":3000")
}
