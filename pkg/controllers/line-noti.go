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
		if payload.Status == "firing" {
			payload.Status = "🔥" + payload.Status + "🔥"
		} else if payload.Status == "resolved" {
			payload.Status = "👌" + payload.Status + "👌"
		}
		go func(alert models.Alerts) {
			message := fmt.Sprintf("\nStatus: %s\nAlertname: %s\nSummary: %s\nDesc: %s",
				payload.Status, alert.Labels.AlertName, alert.Annotations.Summary, alert.Annotations.Description)

			requester.PostForm("https://notify-api.line.me/api/notify", message, token)
		}(alert)
	}
}
