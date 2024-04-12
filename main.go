package main

import (
	"alertmanager-gateway/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.POST("/line-notify", controllers.LineNoti)
	r.POST("/msteams-notify", controllers.MsTeamsNoti)

	r.Run(":8080")
}
