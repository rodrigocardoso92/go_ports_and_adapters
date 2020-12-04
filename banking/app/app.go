package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCostumers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
