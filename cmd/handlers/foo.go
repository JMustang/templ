package handlers

import (
	"fmt"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Hello, Foo!"))
	return fmt.Errorf("error")
}
