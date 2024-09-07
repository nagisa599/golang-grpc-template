package handler

import (
	"context"

	todov1 "github.com/nagisa599/golang-grpc-template/gen/go/v1/todo"
	"github.com/nagisa599/golang-grpc-template/internal/domain"
)
type TodoUsecase interface {
	GetTodo(ctx context.Context,id int) (*domain.Todo, error)
	CreateTodo(ctx context.Context,todo *domain.Todo) error
}
type TodoHandler struct {
	todov1.UnimplementedTodoServiceServer
	tu TodoUsecase
}

func NewTodoHandler(todoUsecase TodoUsecase) *TodoHandler {
	return &TodoHandler{tu: todoUsecase}
}

func (th *TodoHandler) GetTodo(ctx context.Context, req *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	todo,err := th.tu.GetTodo(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &todov1.GetTodoResponse{
		Id: int32(todo.ID),
		Title: todo.Title,
		Description: todo.Description,
	}, nil
}
func (th *TodoHandler) CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	todo := &domain.Todo {
		Title: req.Title,
		Description: req.Description,
	}
	if err := th.tu.CreateTodo(ctx, todo); err != nil {
		return nil, err
	}
	return &todov1.CreateTodoResponse{
		Status: todov1.Status_SUCCESS, // Enum値の参照方法
	}, nil
}