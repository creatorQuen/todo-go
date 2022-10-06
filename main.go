package main

import (
	"context"
	"github.com/gin-gonic/gin"
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

	//r := mux.NewRouter()
	//r.HandleFunc("/todos", handlers.SearchTodos(client)).Methods("GET")
	//r.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods("GET")
	//r.HandleFunc("/todos", handlers.InsertTodo(client)).Methods("POST")
	//r.HandleFunc("/todos/{id}", handlers.UpdateTodo(client)).Methods("PATCH")
	//r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods("DELETE")
	//http.ListenAndServe(":8080", r)

	r := gin.Default()
	todos := r.Group("/todos")
	todos.Use(Authorization(conf.Token))
	{
		todos.GET("/", handlers.SearchTodos(client))
		todos.GET("/:id", handlers.GetTodo(client))
		todos.POST("/", handlers.InsertTodo(client))
		todos.PATCH("/:id", handlers.UpdateTodo(client))
		todos.DELETE("/:id", handlers.DeleteTodo(client))
	}

	r.POST("graphql", handlers.GraphqlTodos(client))

	r.Run(":8080")
}

func Authorization(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if token != auth {
			c.AbortWithStatusJSON(4401, gin.H{"message": "Invalid authorization token"})
		}
		c.Next()
	}
}
