package main

import (
	"github.com/go-chi/chi"
	handlers "main.go/internal"
	"net/http"
)

type App struct {
	router *chi.Mux
}

func NewApp() *App {
	return &App{
		router: chi.NewRouter(),
	}
}

func (a *App) run() {
	//a.router.Post("/user", createUser.Create)
	a.router.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.Create) // POST /users
		r.Route("/{userId}", func(r chi.Router) {
			r.Post("/make_friends", handlers.MakeFriends)
			r.Delete("/", handlers.Delete)
			r.Get("/friends", handlers.GetFriends)
			r.Patch("/age", handlers.UpdateAge)
		})
	})
	http.ListenAndServe(":3333", a.router)

}

func main() {
	var app = NewApp()
	app.run()
}
