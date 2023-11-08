package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/leather-goods", app.listLeatherGoodsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/leather-goods", app.createLeatherGoodsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/leather-goods/:id", app.showLeatherGoodsHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/leather-goods/:id", app.updateLeatherGoodsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/leather-goods/:id", app.deleteLeatherGoodsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	return app.recoverPanic(app.rateLimit(router))

}
