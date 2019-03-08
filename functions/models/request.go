package models

import (
	"time"
)

type Request struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Slots     int       `json:"slots"`
	RideID    string    `json:"rideID"`
	UserID    string    `json:"userID"`
	Status    string    `json:"status"`
}