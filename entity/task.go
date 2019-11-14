package entity

import (
	"errors"
	"time"

	"github.com/grandcolline/todo-list-api/entity/task"
)

// Task タスクエンティティ
type Task struct {
	ID          task.ID          // タスクID
	Name        task.Name        // タスク名
	Description task.Description // タスク詳細
	Status      task.Status      // タスクステータス
	CreatedAt   time.Time        // 作成日時
	UpdatedAt   time.Time        // 更新日時
}

// NewTask はタスクエンティティの初期化をする
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

// Update はタスクエンティティを更新する
func (t *Task) Update(name task.Name, des task.Description) {
	t.Name = name
	t.Description = des
	t.CreatedAt = time.Now()
	return
}

// Complate はタスクエンティティを完了にする
func (t *Task) Complate() error {
	if t.Status == task.Complate {
		return errors.New("complated")
	}
	t.Status = task.Complate
	t.UpdatedAt = time.Now()
	return nil
}
