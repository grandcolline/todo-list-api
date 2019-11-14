package repository

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
)

// TaskRepository タスクレポジトリ
type TaskRepository interface {
	Create(*entity.Task) error
	ReadByID(task.ID) (*entity.Task, error)
	Update(*entity.Task) error
	Delete(*entity.Task) error
	IsNotFound(error) bool
}
