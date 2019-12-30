package entity

import (
	"time"

	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/util/errors"
)

// Task TaskEntity
type Task struct {
	ID          task.ID          // タスクID
	Name        task.Name        // タスク名
	Description task.Description // タスク詳細
	Status      task.Status      // タスクステータス
	CreatedAt   time.Time        // 作成日時
	UpdatedAt   time.Time        // 更新日時
}

// NewTask はタスクの初期化をする
func NewTask() *Task {
	return &Task{
		ID:          task.NewID(),
		Name:        task.NewName(),
		Description: task.NewDescription(),
		Status:      task.NewStatus(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Update はタスクを更新する
func (t *Task) Update(name task.Name, des task.Description) {
	t.Name = name
	t.Description = des
	t.UpdatedAt = time.Now()
}

// Complate はタスクを完了にする
func (t *Task) Complate() error {
	// doingのものしか完了にできない
	if !t.Status.IsDoing() {
		return errors.New(errors.BadParams, "failed to completing task: status is not doing")
	}

	t.Status = task.Complate
	t.UpdatedAt = time.Now()
	return nil
}
