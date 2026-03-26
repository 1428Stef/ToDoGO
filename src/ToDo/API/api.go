package main

import (
	"ToDo/API/endpoint"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/list", endpoint.ListHandler)
	http.HandleFunc("/add", endpoint.AddHandler)
	http.HandleFunc("/done", endpoint.DoneHandler)
	http.HandleFunc("/del", endpoint.DelHandler)

	fmt.Println("Starting HTTP server!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}