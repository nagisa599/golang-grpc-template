package usecase

import (
	"context"

	"github.com/nagisa599/golang-grpc-template/internal/domain"
)
type TodoRepository interface{
	GetTodo(ctx context.Context,id int) (*domain.Todo, error)
	CreateTodo(ctx context.Context,todo *domain.Todo)  error
}

type TodoUsecase struct {
	tr TodoRepository
}

func NewTodoUsecase(todoRepository TodoRepository) *TodoUsecase {
	return &TodoUsecase{tr: todoRepository}
}
func (tu *TodoUsecase) GetTodo(ctx context.Context,id int) (*domain.Todo, error) {
	todo,err:= tu.tr.GetTodo(ctx,id)
	if err != nil {
		return nil,err
	}
	return todo,nil
}
func (tu *TodoUsecase) CreateTodo(ctx context.Context,todo *domain.Todo) error {
	if err := tu.tr.CreateTodo(ctx,todo); err != nil {
		return err
	}
	return nil	
}
