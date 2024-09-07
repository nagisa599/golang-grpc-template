package repository

import (
	"context"

	"github.com/nagisa599/golang-grpc-template/internal/domain"
)

type TodoEntity struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
}

type TodoRepository struct {
	dh DatabaseHandler
}

func NewTodoRepository(databaseHandler DatabaseHandler) *TodoRepository {
	return &TodoRepository{dh: databaseHandler}
}
func newTodoEntity(todo domain.Todo) *TodoEntity {
	return &TodoEntity{
		ID:          uint(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
	}
}
func (tr *TodoRepository) GetTodo(ctx context.Context,id int) (*domain.Todo, error) {
	todoEntity := &TodoEntity{}
	if err := tr.dh.Conn(ctx).Table("todos").Where("id = ?", id).First(todoEntity).Error; err != nil {
		return nil, err
	}
	return &domain.Todo{
		ID:          int(todoEntity.ID),
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
	}, nil
}
func (tr *TodoRepository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	newTodo := newTodoEntity(*todo)
	if err := tr.dh.Conn(ctx).Table("todos").Create(newTodo).Error; err != nil {
		return err
	}
	return nil
}