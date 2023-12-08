package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*file", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreateForm)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	router.HandlerFunc(http.MethodGet, "/user/signup", app.userSignupForm)
	router.HandlerFunc(http.MethodPost, "/user/signup", app.userSignupPost)
	router.HandlerFunc(http.MethodGet, "/user/login", app.userLoginForm)
	router.HandlerFunc(http.MethodPost, "/user/login", app.userLoginPost)
	router.HandlerFunc(http.MethodPost, "/user/logout", app.userLogoutPost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders, app.sessionManager.LoadAndSave)
	return standard.Then(router)
}
