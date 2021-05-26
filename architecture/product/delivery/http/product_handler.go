package http

import (
	"encoding/json"
	"github.com/Imtiaz-Riton/auth"
	"github.com/Imtiaz-Riton/domain/entity"
	"github.com/Imtiaz-Riton/domain/exception"
	"github.com/Imtiaz-Riton/domain/repository"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type productHandler struct {
	repository repository.ProductRepository
}

func (handler productHandler) fetchAll(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	products, err := handler.repository.GetAll()
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}
	if err := json.NewEncoder(rw).Encode(products); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (handler productHandler) fetchByID(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	key := vars["id"]

	product, err := handler.repository.GetByID(key)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(rw).Encode(product); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (handler productHandler) fetchByTitle(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	key := vars["title"]

	product, err := handler.repository.GetByTitle(key)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}
	if err := json.NewEncoder(rw).Encode(product); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (handler productHandler) addNewProduct(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var newProduct *entity.Product
	if err := json.Unmarshal(reqBody, &newProduct); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	product, err := handler.repository.Create(newProduct)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}
	if err := json.NewEncoder(rw).Encode(product); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler productHandler) updateProduct(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(req)
	key := vars["id"]

	var newProduct *entity.Product
	if err := json.Unmarshal(reqBody, &newProduct); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := handler.repository.Update(key, newProduct)
	if err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}

	if err := json.NewEncoder(rw).Encode(product); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (handler productHandler) deleteProduct(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	if err := handler.repository.Delete(key); err != nil {
		errCode := exception.GetStatusCode(err)
		http.Error(rw, err.Error(), errCode)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func NewProductHandler(router *mux.Router, productRepository repository.ProductRepository)  {
	handler := &productHandler{repository: productRepository}

	r := router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/products", auth.JwtAuthentication(handler.fetchAll)).Methods("GET")
	r.HandleFunc("/product/{id}", auth.JwtAuthentication(handler.fetchByID)).Methods("GET")
	r.HandleFunc("/product/{id}", auth.JwtAuthentication(handler.fetchByTitle)).Methods("GET")
	r.HandleFunc("/product", auth.JwtAuthentication(handler.addNewProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", auth.JwtAuthentication(handler.updateProduct)).Methods("PUT")
	r.HandleFunc("/product/{id}", auth.JwtAuthentication(handler.deleteProduct)).Methods("DELETE")
}
