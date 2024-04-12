package controllers

import (
	"alertmanager-gateway/pkg/models"
	"alertmanager-gateway/pkg/requester"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LineNoti(c *gin.Context) {
	token := c.Query("token")

	payload := &models.AlrtMngReq{}
	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})
	for _, alert := range payload.Alerts {
		go func(alert models.Alerts) {
			message := fmt.Sprintf("\nStatus: %s\nAlertname: %s\nMessage: %s",
				payload.Status, alert.Labels.AlertName, alert.Annotations.Summary)

			requester.PostForm("https://notify-api.line.me/api/notify", message, token)
		}(alert)
	}
}
