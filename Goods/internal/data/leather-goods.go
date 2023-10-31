package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
INSERT INTO leatherGoods (name, type, price, leather_type, color)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, version`

	args := []interface{}{leatherGoods.Name, leatherGoods.Type, leatherGoods.Price, leatherGoods.LeatherType, leatherGoods.Color}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return l.DB.QueryRowContext(ctx, query, args...).Scan(&leatherGoods.ID, &leatherGoods.CreatedAt, &leatherGoods.Version)
}

type LeatherGoodsModel struct {
	DB *sql.DB
}

func (l LeatherGoodsModel) Get(id int64) (*LeatherGoods, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, name, type, price, leather_type, color, version
		FROM leatherGoods
		WHERE id = $1`

	var leatherGoods LeatherGoods
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := l.DB.QueryRowContext(ctx, query, id).Scan(
		&leatherGoods.ID,
		&leatherGoods.CreatedAt,
		&leatherGoods.Name,
		&leatherGoods.Type,
		&leatherGoods.Price,
		&leatherGoods.LeatherType,
		&leatherGoods.Color,
		&leatherGoods.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &leatherGoods, nil
}

func (l LeatherGoodsModel) Update(leatherGoods *LeatherGoods) error {
	query := `
		UPDATE leatherGoods
		SET name = $1, type = $2, price = $3, leather_type = $4, color = $5, version = version + 1
		WHERE id = $6 AND version = $7
		RETURNING version`

	args := []interface{}{
		leatherGoods.Name,
		leatherGoods.Type,
		leatherGoods.Price,
		leatherGoods.LeatherType,
		leatherGoods.Color,
		leatherGoods.ID,
		leatherGoods.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := l.DB.QueryRowContext(ctx, query, args...).Scan(&leatherGoods.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (l LeatherGoodsModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM leatherGoods
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := l.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
func (l LeatherGoodsModel) GetAll(name string, color string, filters Filters) ([]*LeatherGoods, error) {

	query := `
		SELECT id, created_at, name, type, price, leather_type, color, version
		FROM leatherGoods
		ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := l.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	leathergoods := []*LeatherGoods{}

	for rows.Next() {

		var leatherGoods LeatherGoods

		err := rows.Scan(
			&leatherGoods.ID,
			&leatherGoods.CreatedAt,
			&leatherGoods.Name,
			&leatherGoods.Type,
			&leatherGoods.Price,
			&leatherGoods.LeatherType,
			&leatherGoods.Color,
			&leatherGoods.Version,
		)
		if err != nil {
			return nil, err
		}

		leathergoods = append(leathergoods, &leatherGoods)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return leathergoods, nil
}
