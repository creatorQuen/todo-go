package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-go/models"
)

func AddNewTodo() string {
	todo := models.Todo{
		UserID:    1,
		Title:     "buy carrot",
		Completed: false,
	}
	res, _ := client.Insert(todo)
	return res.ID.(primitive.ObjectID).Hex()
}

func TestUpdateTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id            string
		payload       string
		expectedCode  int
		modifiedCount int64
	}{
		"should return 200 and modified count 1": {
			id:            id,
			payload:       `{"completed": true}`,
			expectedCode:  200,
			modifiedCount: 1,
		},
		"should return 200 modified count 0": {
			id:            id,
			payload:       `{"title": "buy carrot"}`,
			expectedCode:  200,
			modifiedCount: 0,
		},
		"should return 400": {
			id:           id,
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("PATCH", "/todos/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", UpdateTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.TodoUpdate{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.modifiedCount, todo.ModifiedCount)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	_, _ = client.Delete(id)
}
