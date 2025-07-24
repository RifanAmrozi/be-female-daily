package model

type Mission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	Status      string `json:"status"` // e.g., "Pending", "Completed"
	Brand       Brand  `json:"brand"`  // Associated brand for the mission
}
