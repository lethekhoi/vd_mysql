package product_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"vd_mysql/config"
	"vd_mysql/entities"
	"vd_mysql/models"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}
func SearchPrice(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	smin := vars["min"]
	min, _ := strconv.ParseFloat(smin, 64)
	smax := vars["max"]
	max, _ := strconv.ParseFloat(smax, 64)
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchPrice(min, max)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}
func SearchID(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(path.Base(request.URL.Path))
	if err != nil {
		return
	}
	fmt.Println("url", request.URL.Path)
	fmt.Println("id", id)
	vars := mux.Vars(request)
	sID := vars["ID"]
	ID, _ := strconv.ParseInt(sID, 10, 64)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchIDRow(ID)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}
func Create(response http.ResponseWriter, request *http.Request) {

	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.CreateWithPrepared(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func Update(response http.ResponseWriter, request *http.Request) {

	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sID := vars["ID"]
	ID, _ := strconv.ParseInt(sID, 10, 64)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		RowsAffected, err2 := productModel.Delete(ID)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"RowAffected": RowsAffected,
			})
		}
	}
}

func Count(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	}
	{
		productModel := models.ProductModel{
			Db: db,
		}

		count, err2 := productModel.CountProduct()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]int64{
				"Total Product": count,
			})
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
