package main

import (
	"fmt"
	product_handler"github.com/Imtiaz-Riton/architecture/product/delivery/http"
	"github.com/Imtiaz-Riton/architecture/product/repository/in_memory"
	"github.com/Imtiaz-Riton/auth"
	"github.com/Imtiaz-Riton/domain/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", auth.JwtBasicAuthentication(logIn)).Methods("GET")

	product := entity.Product{
		ID: "12345",
		Title: "Walton",
		Amount: 123,
		Price: 2345,
	}

	productMap := make(map[string]*entity.Product)
	productMap["12"] = &product

	productRepo := in_memory.NewInMemoryProductRepository(productMap)
	product_handler.NewProductHandler(router, productRepo)
	log.Println("listening to port : 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func logIn(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	fmt.Println("Successfully logged in")

	token, err := auth.GenerateJWT()
	fmt.Println("token : ", token)

	if err != nil {
		log.Fatal(err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Token", token)
	rw.Write([]byte(token))
}
