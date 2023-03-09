package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	st "main.go/internal/storage"
	"main.go/internal/user"
	"net/http"
	"strconv"
)

type AnswerCreate struct {
	Id int `json:"id"`
}

type FriendTarget struct {
	Target    int    `json:"target_id,omitempty"`
	Username1 string `json:"username1,omitempty"`
	Username2 string `json:"username2,omitempty"`
}

type Target struct {
	Target int `json:"target_id,omitempty"`
}

type Answer struct {
	Username string `json:"username,omitempty"`
}

type AnswerFriends struct {
	Friends []int `json:"friends"`
}

type NewAge struct {
	Age int `json:"new_age"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create method")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var u user.User
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	newUserId := st.Repository.CreateUser(&u)

	ans := AnswerCreate{newUserId}
	answer, err := json.Marshal(ans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(answer)
}

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("make friend method")
	userId := chi.URLParam(r, "userId")
	uId, err := strconv.Atoi(userId)
	if err != nil || uId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var tar FriendTarget
	if err := json.Unmarshal(content, &tar); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	friends, err := st.Repository.MakeFriends(uId, tar.Target)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Пользователи либо уже дружат, либо не найдены."))
		return
	}
	tar.Username1 = friends[0]
	tar.Username2 = friends[1]
	answer, err := json.Marshal(tar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete user method")
	userId := chi.URLParam(r, "userId")
	uId, err := strconv.Atoi(userId)
	if err != nil || uId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username, err := st.Repository.DeleteUser(uId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var a Answer
	a.Username = username
	answer, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user friends method")
	userIdS := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(userIdS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var friends AnswerFriends
	friendsSlise, err := st.Repository.GetFriends(userId)
	friends.Friends = friendsSlise
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Друзей у пользователя не найдено"))
		return
	}
	result, err := json.Marshal(friends)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}

func UpdateAge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user age method")
	userIdS := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(userIdS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
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
		return
	}
	err = st.Repository.UpdateAge(userId, age.Age)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Пользователя с таким id не существует."))
		return
	}
	w.WriteHeader(http.StatusOK)
}
