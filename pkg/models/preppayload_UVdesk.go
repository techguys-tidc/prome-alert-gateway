package models

import (
	"encoding/json"
	"fmt"
	"os"
)

func GenerateJSON_ToGetAccessToken() string {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	if username == "" || password == "" {
		fmt.Println("Missing environment variables")
		return ""
	}

	UVdeskAccessToken := &UVdeskAccessToken{}
	UVdeskAccessToken.Username = username
	UVdeskAccessToken.Password = password

	jsonData, err := json.MarshalIndent(UVdeskAccessToken, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}
	fmt.Println(string(jsonData))
	return string(jsonData)
}

// ConvertToUVdesk แปลง AlrtMngReq ไปเป็น UVdeskTicket
func ConvertToUVdesk(alert Alerts) UVdeskTicket {
	uvdesk := UVdeskTicket{}
	uvdesk.Name = alert.Labels.AlertName
	uvdesk.Message = alert.Labels.Message
	uvdesk.ActAsType = alert.Labels.ActAsType
	uvdesk.Subject = alert.Labels.Subject
	uvdesk.From = alert.Labels.From
	uvdesk.Values.FirstName = alert.Labels.AlertName
	uvdesk.Values.LastName = "Integration User"
	uvdesk.Values.Description = alert.Labels.Subject
	uvdesk.Values.DetailedDescription = alert.Annotations.Summary
	uvdesk.Values.AffectedCurrentSite = alert.Labels.Instance
	uvdesk.Values.AssignedSupportCompany = "Lotus's"
	uvdesk.Values.AssignedSupportOrganization = "Infrastructure"
	uvdesk.Values.AssignedGroup = "THL3 Service Management ITSM"
	uvdesk.Values.Impact = "2000"
	uvdesk.Values.Urgency = "2000"
	uvdesk.Values.ReportedSource = "Web"
	uvdesk.Values.ServiceType = "Incident"
	uvdesk.Values.ProductCategorizationTier1 = "Integration"
	uvdesk.Values.ProductCategorizationTier2 = "Integration"
	uvdesk.Values.ProductCategorizationTier3 = "NA"
	return uvdesk
}

// GenerateUVdeskJSON แปลง UVdeskTicket เป็น JSON
func GenerateUVdeskJSON(uvdeskTicket UVdeskTicket) string {
	jsonData, err := json.MarshalIndent(uvdeskTicket, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}
	return string(jsonData)
}
