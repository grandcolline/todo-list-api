package usecase

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/usecase/repository"
	"github.com/grandcolline/todo-list-api/util/errors"
)

// User ユーザユースケース
type User struct {
	taskRepo repository.Task
}

// NewUser はユーザユースケースを作成する
func NewUser(taskRepo repository.Task) *User {
	return &User{
		taskRepo: taskRepo,
	}
}

// GetByID はIDでタスクを取得する
func (u *User) GetByID(id task.ID) (*entity.Task, error) {
	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return nil, errors.Errorf("%w", err)
	}

	return task, nil
}

// Create はタスクを登録する
func (u *User) Create(name task.Name, des task.Description) (*entity.Task, error) {
	// タスクを作成
	task := entity.NewTask()
	task.Update(name, des)

	// DB保存
	if err := u.taskRepo.Upsert(task); err != nil {
		return nil, errors.Errorf("%w", err)
	}

	return task, nil
}

// Update はタスクを更新する
func (u *User) Update(id task.ID, name task.Name, des task.Description) error {
	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return errors.Errorf("%w", err)
	}

	// タスク更新
	task.Update(name, des)

	// DB保存
	if err := u.taskRepo.Upsert(task); err != nil {
		return errors.Errorf("%w", err)
	}

	return nil
}

// Complate はタスクを完了にする
func (u *User) Complate(id task.ID) error {
	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return errors.Errorf("%w", err)
	}

	// ステータス更新
	if err = task.Complate(); err != nil {
		return errors.Errorf("%w", err)
	}

	// DB保存
	if err := u.taskRepo.Upsert(task); err != nil {
		return errors.Errorf("%w", err)
	}

	return nil
}
