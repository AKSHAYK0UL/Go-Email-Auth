package models

type Email struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	DeviceToken string `json:"devicetoken"`
	Vcode       string `json:"vcode"`
	ToEmail     string `json:"toemail"`
	FromEmail   string `json:"fromemail"`
	AppPassword string `json:"apppassword"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
