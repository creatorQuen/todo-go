package database

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"todo-go/models"
)

type MockTodoClient struct {
	mock.Mock
}

func (m *MockTodoClient) Insert(todo models.Todo) (models.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoClient) Update(id string, update interface{}) (models.TodoUpdate, error) {
	args := m.Called(id, update)
	return args.Get(0).(models.TodoUpdate), args.Error(1)
}

func (m *MockTodoClient) Delete(id string) (models.ToDoDelete, error) {
	args := m.Called(id)
	return args.Get(0).(models.ToDoDelete), args.Error(1)
}

func (m *MockTodoClient) Get(id string) (models.Todo, error) {
	fmt.Println("call get mock function")
	args := m.Called(id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoClient) Search(filter interface{}) ([]models.Todo, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Todo), args.Error(1)
}
