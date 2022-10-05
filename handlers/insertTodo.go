package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-go/database"
	"todo-go/models"
)

func InsertTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		todo := models.Todo{}
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

//func InsertTodo(db database.TodoInterface) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		todo := models.Todo{}
//
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		err = json.Unmarshal(body, &todo)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		res, err := db.Insert(todo)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		WriteResponse(w, http.StatusOK, res)
//	}
//}
//
//func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	json.NewEncoder(w).Encode(res)
//}
