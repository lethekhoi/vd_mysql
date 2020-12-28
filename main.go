package main

import (
	"fmt"
	"net/http"

	"vd_mysql/apis/product_api"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home!\n")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/count", product_api.Count).Methods("GET")
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods("GET")
	router.HandleFunc("/api/product/searchprice/{min}/{max}", product_api.SearchPrice).Methods("GET")
	router.HandleFunc("/api/product/searchid/{ID}", product_api.SearchID).Methods("GET")
	router.HandleFunc("/api/product/create", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product/update", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{ID}", product_api.Delete).Methods("DELETE")
	fmt.Println("Listen port")
	err := http.ListenAndServe(":5000", router)
	if err != nil {

		fmt.Println(err)
	}
}
