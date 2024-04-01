package models

type Email struct {
	UserName    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
	DeviceToken string `json:"devicetoken"`
	Vcode       string `json:"vcode"`
	ToEmail     string `json:"toemail" validate:"email,required"`
	FromEmail   string `json:"fromemail" validate:"email,required"`
	AppPassword string `json:"apppassword" validate:"required"`
	Host        string `json:"host" validate:"required"`
	Port        string `json:"port" validate:"required"`
}
