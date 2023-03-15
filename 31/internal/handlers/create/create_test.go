package create

import (
	"main.go/internal/handlers/create/mocks"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		creator UserCreator
	}
	tests := []struct {
		name string
		args args
		want func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name: "base test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userCreator := mocks.NewUserCreator(t)
			if got := New(tt.args.creator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
