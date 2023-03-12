package create_handler

import (
	"encoding/json"
	"fmt"
	"io"
	http "net/http"

	user "main.go/internal/user"
)

type AnswerCreate struct {
	Id int `json:"id"`
}

type UserCreator interface {
	CreateUser(u *user.User) int
}

func New(creator UserCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		newUserId := creator.CreateUser(&u)

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
}
