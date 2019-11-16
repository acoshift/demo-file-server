package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
