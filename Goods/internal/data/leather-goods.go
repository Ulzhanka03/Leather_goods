package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goods/Goods/internal/validator"
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
func ValidateLeatherGoods(v *validator.Validator, leatherGoods *LeatherGoods) {
	v.Check(leatherGoods.Name != "", "name", "must be provided")
	v.Check(len(leatherGoods.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(leatherGoods.Price != 0, "price", "must be provided")
	v.Check(leatherGoods.Type != "", "type", "must be provided")
	v.Check(leatherGoods.LeatherType != "", "leather_type", "must be provided")
	v.Check(leatherGoods.Color != "", "color", "must be provided")
}
func (l LeatherGoodsModel) Insert(leatherGoods *LeatherGoods) error {

	query := `
INSERT INTO leathergoods (name, type, price, leather_type, color)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, version`

	args := []interface{}{leatherGoods.Name, leatherGoods.Type, leatherGoods.Price, leatherGoods.LeatherType, leatherGoods.Color}

	return l.DB.QueryRow(query, args...).Scan(&leatherGoods.ID, &leatherGoods.CreatedAt, &leatherGoods.Version)
}

type LeatherGoodsModel struct {
	DB *sql.DB
}

func (l LeatherGoodsModel) Get(id int64) (*LeatherGoods, error) {
	return nil, nil
}

func (l LeatherGoodsModel) Update(leatherGoods *LeatherGoods) error {
	return nil
}

func (l LeatherGoodsModel) Delete(id int64) error {
	return nil
}
