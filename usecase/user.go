package usecase

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"

	"github.com/grandcolline/todo-list-api/usecase/logger"
	"github.com/grandcolline/todo-list-api/usecase/repository"
)

// User ユーザユースケース
type User struct {
	taskRepo repository.Task
	log      logger.Logger
}

// NewUser はユーザユースケースを作成する
// FIXME: ここで入れる引数ってポインタの方がいい？
func NewUser(taskRepo repository.Task, log logger.Logger) *User {
	return &User{
		taskRepo: taskRepo,
		log:      log,
	}
}

// GetByID はIDでタスクを取得する
func (u *User) GetByID(id task.ID) (*entity.Task, error) {

	u.log.Debug("GetByID")

	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Create はタスクを登録する
func (u *User) Create(name task.Name, des task.Description) (*entity.Task, error) {
	// タスクを作成
	task := entity.NewTask()
	task.Update(name, des)

	// 永続化
	if err := u.taskRepo.Upsert(task); err != nil {
		return nil, err
	}

	return task, nil
}

// Update はタスクを更新する
func (u *User) Update(id task.ID, name task.Name, des task.Description) error {
	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return err
	}

	// タスク更新
	task.Update(name, des)

	// 永続化
	if err := u.taskRepo.Upsert(task); err != nil {
		return err
	}

	return nil
}

// Complate はタスクを完了にする
func (u *User) Complate(id task.ID) error {
	// タスクを取得
	task, err := u.taskRepo.ReadByID(id)
	if err != nil {
		return err
	}

	// ステータス更新
	if err = task.Complate(); err != nil {
		return err
	}

	// 永続化
	if err := u.taskRepo.Upsert(task); err != nil {
		return err
	}

	return nil
}
