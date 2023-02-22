package storage

import (
	"errors"
	"main.go/internal/user"
	"sync"
)

type Storage interface {
	CreateUser(*user.User)
	UpdateAge()
	DeleteUser()
	GetFriends()
	MakeFriends(user1 int, user2 int) [2]string
}

type MainSotrage struct {
	Store map[int]*user.User
	sync.Mutex
}

func (ms *MainSotrage) GetFriends(userId int) (friends []int, err error) {
	for key, val := range ms.Store {
		if key == userId {
			friends = val.Friend
			return
		}
	}
	err = errors.New("Пользователь не найден")
	return nil, err
}

func (ms *MainSotrage) MakeFriends(user1 int, user2 int) (friends [2]string, err error) {
	ms.Lock()
	defer ms.Unlock()
	for key := range ms.Store {
		if key == user1 {
			ms.Store[key].Friend = append(ms.Store[key].Friend, user2)
			friends[0] = ms.Store[key].Name
		} else if key == user2 {
			ms.Store[key].Friend = append(ms.Store[key].Friend, user2)
			friends[1] = ms.Store[key].Name
		}
		if friends[0] != "" && friends[1] != "" {
			return
		}
	}
	err = errors.New("неверный идентификатор")
	return
}

func (ms *MainSotrage) CreateUser(u *user.User) int {
	ms.Lock()
	newUserId := len(ms.Store)
	ms.Store[newUserId] = u
	ms.Unlock()
	return newUserId
}

func (ms *MainSotrage) UpdateAge(userId int, newAge int) error {
	ms.Lock()
	defer ms.Unlock()
	for key, val := range ms.Store {
		if key == userId {
			val.Age = newAge
			return nil
		}

	}
	return errors.New("пользователь не найден")
}

func (ms *MainSotrage) DeleteUser(uId int) (username string, err error) {
	ms.Lock()
	defer ms.Unlock()
	for key, val := range ms.Store {
		if key == uId {
			username = val.Name
			delete(ms.Store, key)
			return
		}
	}
	err = errors.New("Пользователь не найден")
	return
}

func NewStore() *MainSotrage {
	return &MainSotrage{
		Store: make(map[int]*user.User),
	}
}

var Repository = NewStore()
