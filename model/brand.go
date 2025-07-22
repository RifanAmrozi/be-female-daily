package model

type Brand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CrowdCount  int    `json:"crowd_count"`
	Description string `json:"description"`
}
