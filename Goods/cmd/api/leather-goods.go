package main

import (
	"fmt"
	"goods/Goods/internal/data"
	"goods/Goods/internal/validator"
	"net/http"
	"time"
)

func (app *application) createLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string  `json:"name"`
		Type        string  `json:"type"`
		Price       float64 `json:"price"`
		LeatherType string  `json:"leather_type"`
		Color       string  `json:"color"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	leatherGoods := &data.LeatherGoods{
		Name:        input.Name,
		Type:        input.Type,
		Price:       input.Price,
		LeatherType: input.LeatherType,
		Color:       input.Color,
	}
	v := validator.New()
	if data.ValidateLeatherGoods(v, leatherGoods); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.LeatherGoods.Insert(leatherGoods)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/leather-goods/%d", leatherGoods.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"leather-good": leatherGoods}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(input.Price != 0, "price", "must be provided")
	v.Check(input.Type != "", "type", "must be provided")
	v.Check(input.LeatherType != "", "leather_type", "must be provided")
	v.Check(input.Color != "", "color", "must be provided")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

}
func (app *application) showLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	leatherGood := data.LeatherGoods{
		ID:          id,
		CreatedAt:   time.Now(),
		Name:        "Handcrafted Leather Bag",
		Type:        "Portfolio",
		Price:       99.99,
		LeatherType: "Natural",
		Color:       "Brown",
		Version:     1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"leatherGood": leatherGood}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
