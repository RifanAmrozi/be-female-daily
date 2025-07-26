package model

type User struct {
	ID              int      `json:"id"`
	Username        string   `json:"username"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	Point           int      `json:"point"`
	Personalization []string `json:"personalization,omitempty"` // Optional field for user personalization
}
