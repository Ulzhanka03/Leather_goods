package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type LeatherGoods struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Price       float64   `json:"price"`
	LeatherType string    `json:"leather_type"`
	Color       string    `json:"color"`
	Version     int32     `json:"version"`
}

func (l LeatherGoods) MarshalJSON() ([]byte, error) {
	var price string
	if l.Price != 0 {
		price = fmt.Sprintf("%.2f $", l.Price)
	}

	type LeatherGoodsAlias LeatherGoods

	aux := struct {
		LeatherGoodsAlias
		Price string `json:"price"`
	}{
		LeatherGoodsAlias: LeatherGoodsAlias(l),
		Price:             price,
	}
	return json.Marshal(aux)
}
