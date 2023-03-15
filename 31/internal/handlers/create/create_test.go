package create

import (
	"main.go/internal/handlers/create/mocks"
	"main.go/internal/user"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		creator UserCreator
		user    user.User
	}
	tests := []struct {
		name string
		args args
		want func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name: "base test",
			args: args{
				user: user.User{
					Name:   "lalala",
					Age:    66,
					Friend: []int{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userCreator := mocks.NewUserCreator(t)
			userCreator.On("CreateUser", tt.args.user).Return(20)
			if got := New(tt.args.creator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
