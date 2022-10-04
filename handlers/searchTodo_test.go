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

func TestSearchTodos(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		payload      string
		expectedCode int
		expected     string
	}{
		"should return 200 - found": {
			payload:      `{"title":"buy carrot"}`,
			expectedCode: 200,
			expected:     "buy carrot",
		},
		"should return 200 - not found": {
			payload:      `{"title":"buy jeans"}`,
			expectedCode: 200,
			expected:     "",
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/todos?q="+test.payload, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos", SearchTodos(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todos := []models.Todo{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todos)
				for _, todo := range todos {
					assert.Equal(t, test.expected, todo.Title)
				}
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

	_, _ = client.Delete(id)
}
