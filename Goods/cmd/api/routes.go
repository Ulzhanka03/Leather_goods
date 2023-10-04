package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/leather-goods", app.listLeatherGoodsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/leather-goods/:id", app.showLeatherGoodHandler)
	router.HandlerFunc(http.MethodPost, "/v1/leather-goods", app.createLeatherGoodHandler)
	router.HandlerFunc(http.MethodPut, "/v1/leather-goods/:id", app.updateLeatherGoodHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/leather-goods/:id", app.deleteLeatherGoodHandler)

	return router
}
