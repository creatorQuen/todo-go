package handlers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/database"
	"todo-go/models"
)

func TestDeleteTodo(t *testing.T) {
	//id := AddNewTodo()
	//
	//tests := map[string]struct {
	//	id           string
	//	expectedCode int
	//	deletedCount int64
	//}{
	//	"should return 200 and deleted count 1": {
	//		id:           id,
	//		expectedCode: 200,
	//		deletedCount: 1,
	//	},
	//	"should return 200 deleted count 0": {
	//		id:           id,
	//		expectedCode: 200,
	//		deletedCount: 0,
	//	},
	//	"should return 400": {
	//		id:           "another-id",
	//		expectedCode: 400,
	//	},
	//	"should return 404 id is empty": {
	//		id:           "",
	//		expectedCode: 404,
	//	},
	//}
	//
	//for name, test := range tests {
	//	t.Run(name, func(t *testing.T) {
	//		req, _ := http.NewRequest("DELETE", "/todos/"+test.id, nil)
	//		rec := httptest.NewRecorder()
	//
	//		r := mux.NewRouter()
	//		r.HandleFunc("/todos/{id}", DeleteTodo(client))
	//		r.ServeHTTP(rec, req)
	//
	//		if test.expectedCode == 200 {
	//			todo := models.ToDoDelete{}
	//			_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
	//			assert.Equal(t, test.deletedCount, todo.DeletedCount)
	//		}
	//
	//		assert.Equal(t, test.expectedCode, rec.Code)
	//	})
	//}
	//_, _ = client.Delete(id)

	//-------------------------------------------------

	client := &database.MockTodoClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		expectedCode int
		expected     string
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
		"should return 404 id is empty": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Delete", test.id).Return(models.ToDoDelete{}, nil) //---
			}

			req, _ := http.NewRequest("DELETE", "/todos/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", DeleteTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Delete")
			}
		})
	}
}
