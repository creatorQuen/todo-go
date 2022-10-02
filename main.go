package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"todo-go/config"
	"todo-go/database"
	"todo-go/handlers"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.TodoClient{
		Ctx: ctx,
		Col: collection,
	}

	r := mux.NewRouter()

	r.HandleFunc("/todos", handlers.SearchTodos(client)).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods("GET")
	r.HandleFunc("/todos", handlers.InsertTodo(client)).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo(client)).Methods("PATCH")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
