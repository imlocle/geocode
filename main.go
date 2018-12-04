package main

import (
    "encoding/json"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"fmt"
)

type Point struct{
	Id			string 	`json: "id,omitempty"`
	Address1	string	`json: "address1,omitempty"`
	Address2	string 	`json: "address2,omitempty"`
	City		string 	`json: "city,omitempty"`
	Province	string 	`json: "province,omitempty"`
	PostalCode	string 	`json: "postalcode,omitempty"`
	Latitude	string 	`json: "latitude,omitempty"`
	Longitude	string 	`json: "longitude,omitempty"`
}

var point []Point

func GetPoint(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range point {
		fmt.Println(item.Address1)
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(point)
}

// our main function
func main() {
	router := mux.NewRouter()
	point = append(point, Point{Id: "1", Address1: "20 Pacifica #1000", City: "Irvine", Province: "California", PostalCode: "92618"})
	router.HandleFunc("/point/{id}", GetPoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}