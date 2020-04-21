package repository

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
)

// Task タスクレポジトリ
type Task interface {
	Upsert(*entity.Task) error
	ReadByID(task.ID) (*entity.Task, error)
	Delete(*entity.Task) error
}
