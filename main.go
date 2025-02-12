package main

import (
	"alertmanager-gateway/pkg/controllers"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	username := "sonarqubetest"
	password := "sonarqubetest"
	fmt.Println("print: %s %s", username, password)
	r := gin.New()

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_BODY"))

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
	r.POST("/uvdesk-notify", controllers.UVdeskNoti)

	r.Run(":8080")
}
