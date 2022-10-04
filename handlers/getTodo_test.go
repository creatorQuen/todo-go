package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/models"
)

func TestGetTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id           string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
			expected:     "buy carrot",
		},
		"should return 400": {
			id:           "another-id",
			expectedCode: 400,
		},
		"should return 404 id is empty": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/todos/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", GetTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.Todo{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.expected, todo.Title)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	_, _ = client.Delete(id)
}
