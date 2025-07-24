package model

type DataStorage struct {
	Brands   []Brand   `json:"brands"`
	Missions []Mission `json:"missions"`
	Tickets  []Ticket  `json:"tickets"`
	Users    []User    `json:"users"`
}
