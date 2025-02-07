package controllers

import (
	"alertmanager-gateway/pkg/models"
	"alertmanager-gateway/pkg/requester"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func UVdeskNoti(c *gin.Context) {
	// err := godotenv.Load("../prome-alert-gateway_dev/.env")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file:", err)
		return
	}
	create_ticket_endpoint := os.Getenv("CREATE_TICKET_ENDPOINT")
	access_token_endpoint := os.Getenv("GET_ACCESS_TOKEN_ENDPOINT")

	msgBody := models.GenerateJSON_ToGetAccessToken()
	if create_ticket_endpoint == "" || access_token_endpoint == "" || msgBody == "" {
		fmt.Println("Missing environment variables")
		return
	}

	token := requester.PostUVdeskAccesstoken(access_token_endpoint, msgBody)
	payload := &models.AlrtMngReq{}
	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})

	for _, alert := range payload.Alerts {
		go func(alert models.Alerts) {
			UVdeskMapping := models.ConvertToUVdesk(alert)
			UVdeskTicket := models.GenerateUVdeskJSON(UVdeskMapping)

			fmt.Println(UVdeskTicket)
			requester.PostUVdeskCreateTicket(create_ticket_endpoint, UVdeskTicket, token)
		}(alert)
	}
}
