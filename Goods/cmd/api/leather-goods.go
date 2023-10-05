package main

import (
	"fmt"
	"net/http"
)

func (app *application) createLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new leather good")
}
func (app *application) showLeatherGoodsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of leather good %d\n", id)
}
