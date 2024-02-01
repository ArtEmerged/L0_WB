package models

import "time"

type Item struct {
	OrderUID    string `json:"-" db:"order_uid"`
	ChrtID      uint   `json:"chrt_id" db:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" db:"track_number" validate:"required"`
	Price       uint   `json:"price" db:"price" validate:"required"`
	Rid         string `json:"rid" db:"rid" validate:"required"`
	Name        string `json:"name" db:"name" validate:"required"`
	Sale        uint8  `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  uint   `json:"total_price" db:"total_price" validate:"required"`
	NmID        uint   `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand" validate:"required"`
	Status      uint16 `json:"status" db:"status" validate:"required"`
}

type Payment struct {
	OrderUID     string `json:"-" db:"order_uid"`
	Transaction  string `json:"transaction" db:"transaction" validate:"required"`
	RequestID    string `json:"request_id" db:"request_id" validate:"required"`
	Currency     string `json:"currency" db:"currency" validate:"required"`
	Provider     string `json:"provider" db:"provider" validate:"required"`
	Amount       uint   `json:"amount" db:"amount" validate:"required"`
	PaymentDt    uint   `json:"payment_dt" db:"payment_dt" validate:"required"`
	Bank         string `json:"bank" db:"bank" validate:"required"`
	DeliveryCost uint   `json:"delivery_cost" db:"delivery_cost" validate:"required"`
	GoodsTotal   uint   `json:"goods_total" db:"goods_total" validate:"required"`
	CustomFee    uint   `json:"custom_fee" db:"custom_fee"`
}

type Delivery struct {
	OrderUID string `json:"-" db:"order_uid"`
	Name     string `json:"name" db:"name" validate:"required"`
	Phone    string `json:"phone" db:"phone" validate:"required"`
	ZIP      string `json:"zip" db:"zip" validate:"required"`
	City     string `json:"city" db:"city" validate:"required"`
	Address  string `json:"address" db:"addres" validate:"required"`
	Region   string `json:"region" db:"region" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
}

type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" db:"track_number" validate:"required"`
	Entry             string    `json:"entry" db:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" db:"items" validate:"required"`
	Locale            string    `json:"locale" db:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature" validate:"required"`
	CustomerID        string    `json:"customer_id" db:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" db:"shardkey" validate:"required"`
	SmID              uint      `json:"sm_id" db:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" db:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" validate:"required"`
}
