package handlers

import (
	"fmt"
	"net/http"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get_by_id method")
}
