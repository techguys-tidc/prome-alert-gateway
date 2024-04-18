package main

import (
	"alertmanager-gateway/pkg/controllers"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	debug, _ := strconv.ParseBool("DEBUG_BODY")

	if debug {
		r.Use(func(c *gin.Context) {
			// Read request body
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println("Error reading request body:", err)
			}

			// Log request body
			fmt.Println("Request body:", string(body))

			// Restore the request body to its original state
			c.Request.Body = io.NopCloser(c.Request.Body)

			// Process request
			c.Next()
		})
	}

	r.POST("/line-notify", controllers.LineNoti)
	r.POST("/msteams-notify", controllers.MsTeamsNoti)

	r.Run(":8080")
}
