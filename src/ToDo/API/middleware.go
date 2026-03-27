package main

import "net/http"

func methodOnly(method string, h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }
        h(w, r)
    }
}