package main

import (
	"log"
	"net/http"
)

func main() {
	myfunction := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	}
	http.HandleFunc("/", myfunction)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
