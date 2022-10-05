package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-go/database"
)

func SearchTodos(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter interface{}
		query := c.Query("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

//func SearchTodos(db database.TodoInterface) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var filter interface{}
//		query := r.URL.Query().Get("q")
//
//		if query != "" {
//			err := json.Unmarshal([]byte(query), &filter)
//			if err != nil {
//				WriteResponse(w, http.StatusBadRequest, err.Error())
//				return
//			}
//		}
//
//		res, err := db.Search(filter)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		WriteResponse(w, http.StatusOK, res)
//	}
//}
