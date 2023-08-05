package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"deneme/helper"
)

func main() {
	r := gin.Default()

	// Varlıklar (entities), servisler (services), mantıklar (logics), veritabanları (database) ve güvenlik (security) için yönlendirmeleri ayarlayın
	helper.SetupRoutes(r)

	r.GET("/login", helper.GetAllSecurityEntities)
	// r.GET("/callback", security.HandleCallback)

	fmt.Println("Server is running on http://localhost:8000")

	// Sunucuyu başlatın
	r.Run(":8000")
}
