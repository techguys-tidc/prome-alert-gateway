package models

type AlrtMngReq struct {
	Alerts []Alerts `json:"alerts"`
	Status string   `json:"status"`
}

type Alerts struct {
	Labels       Labels     `json:"labels"`
	Annotations  Annotaions `json:"annotations"`
	GeneratorURL string     `json:"generatorURL"`
}

type Annotaions struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
type Labels struct {
	AlertName string `json:"alertname"`
	Message   string `json:"message"`
	ActAsType string `json:"actAsType"` // uvdesk parameter
	// Name      string `json:"name"`      // uvdesk parameter
	Subject string `json:"subject"` // uvdesk parameter
	From    string `json:"from"`    // uvdesk parameter
}
