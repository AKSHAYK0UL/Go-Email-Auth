package models

type Email struct {
	ToEmail     string `json:"toemail"`
	FromEmail   string `json:"fromemail"`
	AppPassword string `json:"apppassword"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
