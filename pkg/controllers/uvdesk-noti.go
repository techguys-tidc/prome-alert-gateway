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

var ticket_template = `{
    "message": "%s",
    "actAsType": "%s",
    "name": "%s",
    "subject": "%s",
    "from": "%s"
}`

func UVdeskNoti(c *gin.Context) {
	// err := godotenv.Load("../prome-alert-gateway_dev/.env")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	// token := c.Query("token")
	createticket_endpoint := os.Getenv("CREATE_TICKET_ENDPOINT")

	msg, getaccesstoken_endpoint := models.PrepPayload_to_request_AccessTokenUVdesk()
	if msg == "" || getaccesstoken_endpoint == "" {
		fmt.Println("Missing environment variables")
		return
	}

	token := requester.PostGenAccessTokenUVdesk(getaccesstoken_endpoint, msg)

	payload := &models.AlrtMngReq{}

	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})
	for _, alert := range payload.Alerts {
		go func(alert models.Alerts) {
			message := fmt.Sprintf(ticket_template,
				alert.Labels.Message, alert.Labels.ActAsType, alert.Labels.AlertName, alert.Labels.Subject, alert.Labels.From)

			fmt.Println(message)

			requester.PostUVdesk(createticket_endpoint, message, token)
		}(alert)
	}
}
