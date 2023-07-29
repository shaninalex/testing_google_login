package database

type User struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Image     *string  `json:"image"`
	Providers []string `json:"providers"`
}

type RegisterPayload struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
