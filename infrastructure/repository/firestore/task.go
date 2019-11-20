package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/usecase/repository"
)

// TaskRepoImpl タスクレポジトリ実装
type TaskRepoImpl struct {
	Cli *firestore.Client
}

// NewTaskRepoImpl タスクレポジトリ実装を作成する
func NewTaskRepoImpl(cli *firestore.Client) repository.TaskRepository {
	return &TaskRepoImpl{
		Cli: cli,
	}
}

// Create .
func (tr *TaskRepoImpl) Create(*entity.Task) error {
	return nil
}

// ReadByID .
func (tr *TaskRepoImpl) ReadByID(task.ID) (*entity.Task, error) {
	return nil, nil
}

// Update .
func (tr *TaskRepoImpl) Update(*entity.Task) error {
	return nil
}

// Delete .
func (tr *TaskRepoImpl) Delete(*entity.Task) error {
	return nil
}

// IsNotFound .
func (tr *TaskRepoImpl) IsNotFound(error) bool {
	return false
}
