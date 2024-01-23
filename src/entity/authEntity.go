package entity

type UserCredentials struct {
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type RegisterUser struct {
	Username  string `json:"username" binding:"required,min=4"`
	Password  string `json:"password," binding:"required,min=6"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
