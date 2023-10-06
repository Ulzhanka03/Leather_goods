package main

import (
	"fmt"
	"goods/Goods/internal/data"
	"net/http"
	"time"
)

func (app *application) createLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new leather good")
}
func (app *application) showLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//fmt.Fprintf(w, "show the details of leather good %d\n", id)

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
