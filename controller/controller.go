package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammertime1308/Student-REST-API/handler"
)

func StartServer() {
	r := mux.NewRouter()

	fmt.Println("Starting the server...")

	r.HandleFunc("/api/student/{uid}", handler.GetStudent).Methods("GET")
	r.HandleFunc("/api/student/{uid}", handler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/api/student", handler.CreateStudent).Methods("POST")

	http.ListenAndServe(":8080", r)
}
