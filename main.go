package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo-go/config"
	"todo-go/database"
)

func main() {
	conf := config.GetConfig()
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)

}
