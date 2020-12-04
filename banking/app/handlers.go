package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Costumer struct {
	Name    string `json:"fullName" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCostumers(w http.ResponseWriter, r *http.Request) {
	customers := []Costumer{
		{Name: "Rodrigo", City: "Sao Luis", ZipCode: "65065683"},
		{Name: "Isabela", City: "Sao Luis", ZipCode: "65065683"},
		{Name: "Hailton", City: "Sao Luis", ZipCode: "65065683"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
