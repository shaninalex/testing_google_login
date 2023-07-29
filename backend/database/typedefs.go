package database

type User struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Image     string   `json:"image"`
	Providers []string `json:"providers"`
}
