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
	// fmt.Println(string(jsonData))
	return string(jsonData)
}

// ConvertToUVdesk แปลง AlrtMngReq ไปเป็น UVdeskTicket
func ConvertToUVdesk(alert Alerts) UVdeskTicket {
	uvdesk := UVdeskTicket{}
	uvdesk.Name = alert.Labels.AlertName      //test
	uvdesk.Message = alert.Labels.Message     //test
	uvdesk.ActAsType = alert.Labels.ActAsType //test
	uvdesk.Subject = alert.Labels.Subject     //test
	uvdesk.From = alert.Labels.From           //test
	uvdesk.Values.FirstName = alert.Labels.Firstname
	uvdesk.Values.LastName = alert.Labels.Lastname
	uvdesk.Values.Description = alert.Labels.AlertName
	uvdesk.Values.DetailedDescription = ""
	uvdesk.Values.AffectedCurrentSite = alert.Labels.Affectedcurrentsite
	uvdesk.Values.AssignedSupportCompany = alert.Labels.Assignedsupportcompnay
	uvdesk.Values.AssignedSupportOrganization = alert.Labels.Assignedsuupportorganization
	uvdesk.Values.AssignedGroup = alert.Labels.Assignedgroup
	uvdesk.Values.Impact = alert.Labels.Impact
	uvdesk.Values.Urgency = alert.Labels.Urgency
	uvdesk.Values.ReportedSource = alert.Labels.Reportedsource
	uvdesk.Values.ServiceType = alert.Labels.Servicetype
	uvdesk.Values.ProductCategorizationTier1 = alert.Labels.Productcategorizationtier1
	uvdesk.Values.ProductCategorizationTier2 = alert.Labels.Productcategorizationtier2
	uvdesk.Values.ProductCategorizationTier3 = alert.Labels.Productcategorizationtier3
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
