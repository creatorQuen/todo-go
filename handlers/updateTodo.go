package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-go/database"
)

func UpdateTodo(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo interface{}
		id := c.Param("id")
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

//func UpdateTodo(db database.TodoInterface) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		params := mux.Vars(r)
//		id := params["id"]
//
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		var todo interface{}
//		err = json.Unmarshal(body, &todo)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		res, err := db.Update(id, todo)
//		if err != nil {
//			WriteResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		WriteResponse(w, http.StatusOK, res)
//	}
//}
