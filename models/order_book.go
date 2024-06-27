package models

type DepthOrder struct {
	Price   float64
	BaseQty float64
}

type OrderBook struct {
	ID       int64        `json:"id" gorm:"primaryKey"`
	Exchange string       `json:"exchange"`
	Pair     string       `json:"pair"`
	Asks     []DepthOrder `json:"asks" gorm:"serializer:json"`
	Bids     []DepthOrder `json:"bids" gorm:"serializer:json"`
}
