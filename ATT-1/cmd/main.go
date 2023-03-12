package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	city "main.go/internal/city"
	handlers "main.go/internal/handlers"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type App struct {
	router *chi.Mux
	done   chan os.Signal
	store  map[int]city.City
}

func NewApp() *App {
	ret := &App{
		router: chi.NewRouter(),
		done:   make(chan os.Signal, 1),
	}
	signal.Notify(ret.done, os.Interrupt)
	return ret
}

func (a *App) run() {
	//a.router.Post("/user", createUser.Create)
	a.router.Route("/city", func(r chi.Router) {
		r.Post("/", handlers.Create) // POST /users
		r.Route("/{cityId}", func(r chi.Router) {
			r.Get("/", handlers.GetById)
		})
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8000", a.router))
		for {
			select {
			case <-a.done:
				fmt.Println("Exiting")
				return
			}
		}
	}()
	wg.Wait()

}

func main() {
	var app = NewApp()
	app.run()

}
