package main

import (
	"api/cmd/routers"
	"log"
	"net/http"

	"api/cmd/database"

	"github.com/gorilla/mux"

	_ "api/cmd/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	db := database.InitDB("./example.db")
	defer db.Close()
	database.CreateTable(db)

	router := mux.NewRouter()

	router.HandleFunc("/link_create", routers.CreateLink).Methods("POST")
	router.HandleFunc("/{link}", routers.GetLink).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
