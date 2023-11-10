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

	router.HandlerFunc(http.MethodGet, "/v1/leather-goods", app.requireActivatedUser(app.listLeatherGoodsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/leather-goods", app.requireActivatedUser(app.createLeatherGoodsHandler))
	router.HandlerFunc(http.MethodGet, "/v1/leather-goods/:id", app.requireActivatedUser(app.showLeatherGoodsHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/leather-goods/:id", app.requireActivatedUser(app.updateLeatherGoodsHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/leather-goods/:id", app.requireActivatedUser(app.deleteLeatherGoodsHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))

}
