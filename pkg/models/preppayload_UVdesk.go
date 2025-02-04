package models

import (
	"fmt"
	"os"
)

var AccessTokenUVdesk_template = `{
	"username":"%s",
	"password":"%s"
}`

func PrepPayload_to_request_AccessTokenUVdesk() (string, string) {
	// payload := os.Getenv("PAYLOAD")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	url := os.Getenv("GET_ACCESS_TOKEN_ENDPOINT")
	payload := fmt.Sprintf(AccessTokenUVdesk_template, username, password)

	if payload == "" || url == "" {
		fmt.Println("Missing environment variables")
		return "", ""
	}
	return payload, url
}
