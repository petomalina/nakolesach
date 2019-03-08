package models

import (
	"time"
)

type Ride struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Rider           Rider     `json:"rider"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Slots           int       `json:"slots"`
	AvailableSlots  int       `json:"availableSlots"`
	Price           string    `json:"price"`
	Currency        string    `json:"currency"`
	Open            bool      `json:"open"`
	AutoAccept      bool      `json:"autoAccept"`
	BoardingTime    time.Time `json:"boardingTime"`
	BoardingPlace   string    `json:"boardingPlace"`
	DeboardingPlace string    `json:"deboardingPlace"`
}

type Rider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
