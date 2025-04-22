package entity

import "time"

// CreateShipmentRequest holds exactly the fields your clients should send.
type CreateShipmentRequest struct {
	FromAddress     string    `json:"from_address"     validate:"required"`
	ToAddress       string    `json:"to_address"       validate:"required"`
	PickupTime      time.Time `json:"pickup_time"      validate:"required"`
	DeliveryPrice   float64   `json:"delivery_price"   validate:"required"`
	PriceNegotiable bool      `json:"price_negotiable"`
	Weight          float64   `json:"weight"           validate:"required"`
	Volume          float64   `json:"volume"           validate:"required"`
	CargoType       string    `json:"cargo_type"       validate:"required"`
	SenderName      string    `json:"sender_name"      validate:"required"`
	SenderPhone     string    `json:"sender_phone"     validate:"required"`
	ReceiverName    string    `json:"receiver_name"    validate:"required"`
	ReceiverPhone   string    `json:"receiver_phone"   validate:"required"`
	AdditionalNotes string    `json:"additional_notes"`
}
