package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	product "vd_mysql/apis/product_api"
	"vd_mysql/entities"
)

//test Create
// func Test1(t *testing.T) {
// 	new := strings.NewReader(`{"name":"dep nam","price":23,"quantity":15}`)
// 	req, err := http.NewRequest("POST", "/api/product/create", new)
// 	if err != nil {
// 		t.Fatalf("we could not create a request: %v", err)
// 	}
// 	rec := httptest.NewRecorder()
// 	product.Create(rec, req)
// 	res := rec.Result()
// 	defer res.Body.Close()
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK, got %v", res.Status)
// 	}
// 	var product entities.Product
// 	json.Unmarshal(rec.Body.Bytes(), &product)
// 	fmt.Println(product)
// 	if product.Name != "dep nam" {
// 		t.Errorf("Cannot retrieve JSON post")
// 	}
// }

//test FindAll
// func TestFindAll(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/product/findall", nil)
// 	if err != nil {
// 		t.Fatalf("we could not create a request: %v", err)
// 	}
// 	rec := httptest.NewRecorder()
// 	product.FindAll(rec, req)
// 	res := rec.Result()
// 	defer res.Body.Close()
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK, got %v", res.Status)
// 	}
// 	expected := string(`[{"id": 2,"name": "Quan","price": 2,"quantity": 4},{"id": 5,"name": "quan tay","price": 5,"quantity": 12},{"id": 6,"name": "quan kaki","price": 6,"quantity": 10},{"id": 7,"name": "ao somi","price": 17,"quantity": 9},{"id": 22,"name": "dep nu","price": 21,"quantity": 14}]`)
// 	assert.JSONEq(t, expected, rec.Body.String(), "Response body differs")
// }

// //test Search Name
// func TestSearchName(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/product/search/ao", nil)
// 	if err != nil {
// 		t.Fatalf("we could not create a request: %v", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/product/search/{keyword}", product.Search).Methods("GET")
// 	r.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// 	expected := string(`[{"id": 7,"name": "ao somi","price": 17,"quantity": 9}]`)
// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

//test DELETE
// func TestDelete(t *testing.T) {
// 	req, err := http.NewRequest("DELETE", "/api/product/delete/22", nil)
// 	if err != nil {
// 		t.Fatalf("we could not create a request: %v", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/product/delete/{ID}", product.Delete).Methods("DELETE")
// 	r.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// 	expected := string(`{"RowAffected": 1}`)
// 	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
// }

//test SearchID
// func TestSearchID(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/product/searchid/2", nil)
// 	if err != nil {
// 		t.Fatalf("we could not create a request: %v", err)
// 	}
// 	rec := httptest.NewRecorder()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/product/searchid/{ID}", product.SearchID).Methods("GET")
// 	r.ServeHTTP(rec, req)
// 	if status := rec.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
// 	}
// 	var product []entities.Product
// 	json.Unmarshal(rec.Body.Bytes(), &product)
// 	fmt.Println("body", rec.Body)
// 	fmt.Println("product", product)
// 	if len(product) != 1 {
// 		t.Errorf("Expect 1, got %v. ", len(product))
// 	}
// 	if product[0].Name != "Quan" {
// 		t.Errorf("Cannot retrieve JSON post")
// 	}
// }

//testUpdate
func TestUpdate(t *testing.T) {
	item := strings.NewReader(`{"id": 12,"name": "dep da nam","price": 20,"quantity": 11}`)
	req, err := http.NewRequest("PUT", "/api/product/update", item)
	if err != nil {
		t.Fatalf("we could not create a request: %v", err)
	}

	rec := httptest.NewRecorder()
	http.HandlerFunc(product.Update).ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	var product entities.Product

	json.Unmarshal(rec.Body.Bytes(), &product)
	if product.Name != "dep da nam" {
		t.Errorf("Cannot retrieve JSON post")
	}
}
