package entity

import "time"

// Shipment represents the delivery order details mapped to PostgreSQL columns.
type Shipment struct {
	ID              string    `db:"id"`               // unique identifier
	FromAddress     string    `db:"from_address"`     // адрес отправки
	ToAddress       string    `db:"to_address"`       // адрес получения
	PickupTime      time.Time `db:"pickup_time"`      // дата и время забора груза
	DeliveryPrice   float64   `db:"delivery_price"`   // цена доставки
	PriceNegotiable bool      `db:"price_negotiable"` // признак возможности корректировки цены
	Weight          float64   `db:"weight"`           // вес груза (кг)
	Volume          float64   `db:"volume"`           // объём груза (м³)
	CargoType       string    `db:"cargo_type"`       // тип груза
	SenderName      string    `db:"sender_name"`      // имя грузчика на точке отправки
	SenderPhone     string    `db:"sender_phone"`     // номер грузчика на точке отправки
	ReceiverName    string    `db:"receiver_name"`    // имя получателя на точке доставки
	ReceiverPhone   string    `db:"receiver_phone"`   // номер получателя на точке доставки
	AdditionalNotes string    `db:"additional_notes"` // дополнительные комментарии или инструкции
}
