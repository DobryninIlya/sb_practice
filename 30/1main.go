package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Friend []int  `json:"friends"`
}

type service struct {
	store map[int]*User
}

type FriendTarget struct {
	Source    int `json:"source_id"`
	Target    int `json:"target_id"`
	sUsername string
	tUsername string
}
type Target struct {
	Target int `json:"target_id"`
}
type NewAge struct {
	Age int `json:"new_age"`
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(503)
		w.Write([]byte("\nsomething went wrong\n"))
	}
}

func (srv *service) Create(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("create method")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	var u User
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return err
	}
	newUserId := len(srv.store)
	srv.store[newUserId] = &u
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(newUserId)))
	return nil
}

func (srv *service) MakeFriends(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("make friend method")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}
	var tar FriendTarget
	if err := json.Unmarshal(content, &tar); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return err
	}
	var exitFlag uint8
	for key := range srv.store {
		if key == tar.Source {
			srv.store[key].Friend = append(srv.store[key].Friend, tar.Target)
			tar.sUsername = srv.store[key].Name
			exitFlag += 1
		}
		if key == tar.Target {
			srv.store[key].Friend = append(srv.store[key].Friend, tar.Source)
			tar.tUsername = srv.store[key].Name
			exitFlag += 1
		}
		if exitFlag == 2 {
			break
		}
	}
	if exitFlag != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный идентефикатор!"))
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v и %v теперь друзья", tar.sUsername, tar.tUsername)))
	return nil
}

func (srv *service) Delete(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("delete user method")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}
	var tar Target
	if err := json.Unmarshal(content, &tar); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return err
	}
	var username string
	for key, val := range srv.store {
		if key == tar.Target {
			username = val.Name
			delete(srv.store, key)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(username))
			return nil
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Пользователя с таким id не существует."))
	return nil
}

func (srv *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user friends method")
	userIdS := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(userIdS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	var friends []int
	for key, val := range srv.store {
		if key == userId {
			friends = val.Friend
			result, err := json.Marshal(&friends)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(result)
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Пользователя с таким id не существует."))
}

func (srv *service) UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user age method")
	userIdS := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(userIdS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	for key, val := range srv.store {
		if key == userId {
			content, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			var age NewAge
			if err := json.Unmarshal(content, &age); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				fmt.Println(err.Error())
				return
			}
			val.Age = age.Age
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Возраст успешно обновлен!"))
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Пользователя с таким id не существует."))
}

func main() {
	r := chi.NewRouter()
	srv := service{make(map[int]*User)}
	r.Method("POST", "/create", Handler(srv.Create))
	r.Method("POST", "/make_friends", Handler(srv.MakeFriends))
	r.Method("DELETE", "/user", Handler(srv.Delete))
	r.Get("/friends/{userId}", srv.GetFriends)
	r.Put("/{userId}", srv.UpdateUserAge)
	http.ListenAndServe(":3333", r)
}
