package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	handlers "main.go/internal"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

type App struct {
	router *chi.Mux
	done   chan os.Signal
}

func NewApp() *App {
	ret := &App{
		router: chi.NewRouter(),
		done:   make(chan os.Signal, 1),
	}
	signal.Notify(ret.done, os.Interrupt)
	return ret
}

func (a *App) run(port string) {
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":"+port, a.router))
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
	if len(os.Args[1:]) == 0 {
		log.Fatal("Must input port number")
		return
	}
	port := os.Args[1]
	if portNum, err := strconv.Atoi(port); err == nil {
		if portNum > 10000 || portNum < 1000 {
			log.Fatal("Incorrect port number")
			return
		}
		var app = NewApp()
		app.run(port)
	} else {
		log.Fatal("Must input port number")
		return
	}

}
