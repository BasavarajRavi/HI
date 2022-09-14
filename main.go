package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/get", display)
	http.ListenAndServe(":9000", r)
}

func display(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hi hi!"))

	if err != nil {
		panic(err.Error())
	}
}
