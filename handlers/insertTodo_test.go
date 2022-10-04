package handlers

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-go/database"
	"todo-go/models"
)

//var client database.TodoInterface
//
//func init() {
//	conf := config.MongoConfiguration{
//		Server:     "mongodb://localhost:27017",
//		Database:   "Mgo",
//		Collection: "TodosTest",
//	}
//	ctx := context.TODO()
//
//	db := database.ConnectDB(ctx, conf)
//	collection := db.Collection(conf.Collection)
//
//	client = &database.TodoClient{
//		Ctx: ctx,
//		Col: collection,
//	}
//}

func TestInsertTodo(t *testing.T) {
	//tests := map[string]struct {
	//	payload      string
	//	expectedCode int
	//	expected     string
	//}{
	//	"should return 200": {
	//		payload:      `{"userId":1,"title":"buy carrot","completed":false}`,
	//		expectedCode: 200,
	//		expected:     "buy carrot",
	//	},
	//	"should return 400": {
	//		payload:      "invalid json string",
	//		expectedCode: 400,
	//	},
	//}
	//
	//for name, test := range tests {
	//	t.Run(name, func(t *testing.T) {
	//		req, _ := http.NewRequest("POST", "/todos", strings.NewReader(test.payload))
	//		rec := httptest.NewRecorder()
	//		h := InsertTodo(client)
	//		h.ServeHTTP(rec, req)
	//
	//		fmt.Println(rec.Body.String())
	//
	//		if test.expectedCode == 200 {
	//			todo := models.Todo{}
	//			_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
	//			assert.Equal(t, test.expected, todo.Title)
	//			assert.NotNil(t, todo.ID)
	//			// cleanup
	//			_, _ = client.Delete(todo.ID.(string))
	//		}
	//		assert.Equal(t, test.expectedCode, rec.Code)
	//	})
	//}

	//-----------------------------------------------------------

	client := &database.MockTodoClient{}

	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload:      `{"userId":1,"title":"buy carrot","completed":false}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Insert", mock.Anything).Return(models.Todo{}, nil)

			req, _ := http.NewRequest("POST", "/todos", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()
			h := InsertTodo(client)
			h.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
