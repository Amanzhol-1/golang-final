package entity

import "time"

type ShipmentStatus string

const (
	StatusPending   ShipmentStatus = "pending"    // waiting for pickup
	StatusPickedUp  ShipmentStatus = "picked_up"  // cargo is on the move
	StatusInTransit ShipmentStatus = "in_transit" // en route
	StatusDelivered ShipmentStatus = "delivered"  // safely at destination
	StatusCancelled ShipmentStatus = "cancelled"  // shipment call-off
)

// Shipment represents the delivery order details mapped to PostgreSQL columns.
type Shipment struct {
	ID              string         `db:"id"`               // unique identifier
	UserID          string         `db:"user_id"`          // owner (user) of this shipment
	FromAddress     string         `db:"from_address"`     // адрес отправки
	ToAddress       string         `db:"to_address"`       // адрес получения
	PickupTime      time.Time      `db:"pickup_time"`      // дата и время забора груза
	DeliveryPrice   float64        `db:"delivery_price"`   // цена доставки
	PriceNegotiable bool           `db:"price_negotiable"` // признак возможности корректировки цены
	Weight          float64        `db:"weight"`           // вес груза (кг)
	Volume          float64        `db:"volume"`           // объём груза (м³)
	CargoType       string         `db:"cargo_type"`       // тип груза
	SenderName      string         `db:"sender_name"`      // имя грузчика на точке отправки
	SenderPhone     string         `db:"sender_phone"`     // номер грузчика на точке отправки
	ReceiverName    string         `db:"receiver_name"`    // имя получателя на точке доставки
	ReceiverPhone   string         `db:"receiver_phone"`   // номер получателя на точке доставки
	AdditionalNotes string         `db:"additional_notes"` // дополнительные комментарии или инструкции
	Status          ShipmentStatus `db:"status"`
	PickerID        string         `db:"picker_id"`
}
