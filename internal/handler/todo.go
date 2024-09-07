package handler

import (
	"context"

	todov1 "github.com/nagisa599/golang-grpc-template/gen/go/v1/todo"
)

type TodoHandler struct {
	todov1.UnimplementedTodoServiceServer
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (th *TodoHandler) GetTodo(ctx context.Context, req *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	return &todov1.GetTodoResponse{
		Id: 1,
		Title: "title",
		Description: "description",
	}, nil
}