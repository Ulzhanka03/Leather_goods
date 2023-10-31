package main

import (
	"errors"
	"fmt"
	"goods/Goods/internal/data"
	"goods/Goods/internal/validator"
	"net/http"
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
	leatherGood, err := app.models.LeatherGoods.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"leatherGood": leatherGood}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) updateLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	leatherGoods, err := app.models.LeatherGoods.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Name        *string  `json:"name"`
		Type        *string  `json:"type"`
		Price       *float64 `json:"price"`
		LeatherType *string  `json:"leather_type"`
		Color       *string  `json:"color"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Name != nil {
		leatherGoods.Name = *input.Name
	}

	if input.Type != nil {
		leatherGoods.Type = *input.Type
	}
	if input.Price != nil {
		leatherGoods.Price = *input.Price
	}
	if input.LeatherType != nil {
		leatherGoods.LeatherType = *input.LeatherType
	}
	if input.Color != nil {
		leatherGoods.Color = *input.Color
	}

	v := validator.New()
	if data.ValidateLeatherGoods(v, leatherGoods); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.LeatherGoods.Update(leatherGoods)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"leather goods": leatherGoods}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.LeatherGoods.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "leather good successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
