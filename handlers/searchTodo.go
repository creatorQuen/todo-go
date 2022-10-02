package handlers

import (
	"encoding/json"
	"net/http"
	"todo-go/database"
)

func SearchTodos(db database.TodoInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter interface{}
		query := r.URL.Query().Get("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}
