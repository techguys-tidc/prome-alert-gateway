package controllers

import (
	"alertmanager-gateway/pkg/models"
	"alertmanager-gateway/pkg/requester"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var template = `{
	"type":"message",
	"attachments":[
	  {
		"contentType":"application/vnd.microsoft.card.adaptive",
		"contentUrl":null,
		"content":{
		  "$schema":"http://adaptivecards.io/schemas/adaptive-card.json",
		  "type":"AdaptiveCard",
		  "version":"1.0",
		  "body":[
			{
				"type":"TextBlock",
				"text":"**%s**",
				"wrap":true,
				"size":"Medium",
				"spacing":"ExtraLarge"
			},
			{
			  "type":"Container",
			  "padding":"None",
			  "items":[
				%s
			  ]
			}
		  ],
		  "padding":"None"
		}
	  }
	]
  }`

var bodyTemaplte = `{
	"type":"TextBlock",
	"text":"**%s**",
	"wrap":true,
	"size":"Medium",
	"spacing":"ExtraLarge"
  },
  {
	"type":"FactSet",
	"facts":[
	  {
		"title":"Status",
		"value":"***%s***"
	  },
	  {
		"title":"Message",
		"value":"%s"
	  }
	]
  },`

func MsTeamsNoti(c *gin.Context) {
	url := c.Query("url")
	cluster := c.Query("cluster")

	payload := &models.AlrtMngReq{}
	if err := c.BindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})
	contents := ""
	for _, alert := range payload.Alerts {
		contents += fmt.Sprintf(bodyTemaplte,
			alert.Labels.AlertName, payload.Status, alert.Annotations.Summary)
	}

	contents = contents[:len(contents)-1]
	message := fmt.Sprintf(template, cluster, contents)

	requester.PostMsTeams("https://"+url, message)
}
