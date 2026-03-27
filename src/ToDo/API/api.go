package main

import (
	"ToDo/API/endpoint"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/add", methodOnly(http.MethodPost, endpoint.AddHandler))
	http.HandleFunc("/del", methodOnly(http.MethodDelete, endpoint.DelHandler))
	http.HandleFunc("/done", methodOnly(http.MethodPatch, endpoint.DoneHandler))
	http.HandleFunc("/list", methodOnly(http.MethodGet, endpoint.ListHandler))

	fmt.Println("Starting HTTP server!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}