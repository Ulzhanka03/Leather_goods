package data

import "time"

type LeatherGoods struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Price       float64   `json:"price"`
	LeatherType string    `json:"leather_type"`
	Color       string    `json:"color"`
	Version     int32     `json:"version"`
}
