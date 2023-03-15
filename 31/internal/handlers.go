package handlers

import (
	"main.go/internal/user"
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
type UserCreator interface {
	CreateUser(u *user.User) int
}
