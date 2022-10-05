package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-go/database"
)

func GetTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

//func GetTodo(db database.TodoInterface) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		params := mux.Vars(r)
//		id := params["id"]
//
//		res, err := db.Get(id)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		WriteResponse(w, http.StatusOK, res)
//	}
//}
