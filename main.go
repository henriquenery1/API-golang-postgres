package main

import (
	"log"
	"net/http"
)

func main() {
	http.HanleFunc("/", GETHandler)
	http.HandleFunc("/inset", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
