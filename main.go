package main

import (
	"fmt"
	"log"
	"net/http"

	categoryController "./handler/category"
	listController "./handler/list"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/list", listController.ShowAllLists).Methods("GET")
	router.HandleFunc("/list", listController.InsertList).Methods("POST")
	router.HandleFunc("/list/{id:[0-9]+}", listController.UpdateList).Methods("PUT")
	router.HandleFunc("/list/{id:[0-9]+}", listController.DeleteList).Methods("DELETE")

	router.HandleFunc("/category", categoryController.ShowAllCategory).Methods("GET")
	router.HandleFunc("/category", categoryController.InsertCategory).Methods("POST")
	router.HandleFunc("/category/{id:[0-9]+}", categoryController.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id:[0-9]+}", categoryController.DeleteCategory).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")

	// define the port and load routes
	log.Fatal(http.ListenAndServe(":8080", router))

}
