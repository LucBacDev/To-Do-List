package payload

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Avatar string `json:"avatar"`
}