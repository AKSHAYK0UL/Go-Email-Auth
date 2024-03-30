package models

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DeviceToken string `json:"devicetoken"`
	Islogin     bool   `json:"islogin"`
	HostName    string `json:"hostname"`
}
