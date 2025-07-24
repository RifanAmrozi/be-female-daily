package model

type Brand struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CrowdCount    int    `json:"crowd_count"`
	IncreaseCount int    `json:"increase_count"`
	Description   string `json:"description"`
	Category      string `json:"category,omitempty"` // Optional field for brand category
	Hall          string `json:"hall,omitempty"`     // Optional field for brand hall
}
