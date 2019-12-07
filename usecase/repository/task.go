package repository

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
)

// TaskRepository タスクレポジトリ
type TaskRepository interface {
	Upsert(*entity.Task) error
	ReadByID(task.ID) (*entity.Task, error)
	Delete(task.ID) error
}
