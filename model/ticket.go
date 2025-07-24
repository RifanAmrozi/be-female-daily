package model

import "time"

type Ticket struct {
	ID        string    `json:"id"`
	EventName string    `json:"name"`
	Status    string    `json:"status"` // e.g., "Belum ditukar", "Sudah ditukar", "Kadaluarsa"
	ExpiredAt time.Time `json:"expired_at"`
}
