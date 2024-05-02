package dtos

type LoginDTO struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

type RegisterDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
