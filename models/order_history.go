package models

import "time"

type HistoryOrder struct {
	ClientName          string    `json:"client_name" gorm:"column:client_name"`
	ExchangeName        string    `json:"exchange_name" gorm:"column:exchange_name"`
	Label               string    `json:"label" gorm:"column:label"`
	Pair                string    `json:"pair" gorm:"column:pair"`
	Side                string    `json:"side" gorm:"column:side"`
	Type                string    `json:"type" gorm:"column:type"`
	BaseQty             float64   `json:"base_qty" gorm:"column:base_qty"`
	Price               float64   `json:"price" gorm:"column:price"`
	AlgorithmNamePlaced string    `json:"algorithm_name_placed" gorm:"column:algorithm_name_placed"`
	LowestSellPrc       float64   `json:"lowest_sell_prc" gorm:"column:lowest_sell_prc"`
	HighestBuyPrc       float64   `json:"highest_buy_prc" gorm:"column:highest_buy_prc"`
	CommissionQuoteQty  float64   `json:"commission_quote_qty" gorm:"column:commission_quote_qty"`
	TimePlaced          time.Time `json:"time_placed" gorm:"column:time_placed"`
}

type Client struct {
	ClientName   string `json:"client_name" gorm:"column:client_name"`
	ExchangeName string `json:"exchange_name" gorm:"column:exchange_name"`
	Label        string `json:"label" gorm:"column:label"`
	Pair         string `json:"pair" gorm:"column:pair"`
}
