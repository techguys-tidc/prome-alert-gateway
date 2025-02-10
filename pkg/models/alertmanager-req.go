package models

// Variables for handling a message body format from Mimir to PROME-ALERT-GATEWAY
type AlrtMngReq struct {
	Alerts []Alerts `json:"alerts"`
	Status string   `json:"status"`
}

type Alerts struct {
	Labels       Labels     `json:"labels"`
	Annotations  Annotaions `json:"annotations"`
	GeneratorURL string     `json:"generatorURL"`
}

type Labels struct {
	AlertName                    string `json:"alertname"`
	Message                      string `json:"message"`
	ActAsType                    string `json:"actAsType"` // add new mimir parameter
	Subject                      string `json:"subject"`   // add new mimir parameter
	From                         string `json:"from"`      // add new mimir parameter
	Instance                     string `json:"instance"`  // add new mimir parameter
	Firstname                    string `json:"uvdesk_firstname"`
	Lastname                     string `json:"uvdesk_lastname"`
	Affectedcurrentsite          string `json:"uvdesk_affectedcurrentsite"`
	Assignedsupportcompnay       string `json:"uvdesk_assignedsupportcompnay"`
	Assignedsuupportorganization string `json:"uvdesk_assignedsuupportorganization"`
	Assignedgroup                string `json:"uvdesk_assignedgroup"`
	Impact                       string `json:"uvdesk_impact"`
	Urgency                      string `json:"uvdesk_urgency"`
	Reportedsource               string `json:"uvdesk_reportedsource"`
	Servicetype                  string `json:"uvdesk_servicetype"`
	Productcategorizationtier1   string `json:"uvdesk_productcategorizationtier1"`
	Productcategorizationtier2   string `json:"uvdesk_productcategorizationtier2"`
	Productcategorizationtier3   string `json:"uvdesk_productcategorizationtier3"`
}

type Annotaions struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

// Variables for handling a message body format from PROME-ALERT-GATEWAY to UVdesk

// Get Access Token
type UVdeskAccessToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create ticket
type UVdeskTicket struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	ActAsType string `json:"actAsType"`
	Subject   string `json:"subject"`
	From      string `json:"from"`
	Values    Value  `json:"values"`
}

type Value struct {
	FirstName                   string `json:"First_Name"`
	LastName                    string `json:"Last_Name"`
	Description                 string `json:"Description"`
	DetailedDescription         string `json:"Detailed_Decription"`
	AffectedCurrentSite         string `json:"Affected_Current_Site"`
	AssignedSupportCompany      string `json:"Assigned Support Company"`
	AssignedSupportOrganization string `json:"Assigned Support Organization"`
	AssignedGroup               string `json:"Assigned Group"`
	Impact                      string `json:"Impact"`
	Urgency                     string `json:"Urgency"`
	ReportedSource              string `json:"Reported Source"`
	ServiceType                 string `json:"Service_Type"`
	ProductCategorizationTier1  string `json:"Product Categorization Tier 1"`
	ProductCategorizationTier2  string `json:"Product Categorization Tier 2"`
	ProductCategorizationTier3  string `json:"Product Categorization Tier 3"`
}

// UVdesk response body
type UVdeskRespBody struct {
	TicketId int            `json:"ticketId"`
	Value    RespBody_Value `json:"value"`
}

type RespBody_Value struct {
	// response body
	IncidentNumber int `json:"Incident Number"`
}
