package usecase

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"

	// "github.com/grandcolline/todo-list-api/usecase/logger"
	"github.com/grandcolline/todo-list-api/usecase/repository"
)

// TaskUsecase タスクユースケース
type TaskUsecase struct {
	TaskRepo repository.TaskRepository
	// Log      logger.Logger
}

// NewTaskUsecase はタスクユースケースを作成する
func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		TaskRepo: repo,
	}
}

// GetByID はIDでタスクを取得する
func (tu *TaskUsecase) GetByID(id task.ID) (*entity.Task, error) {
	// tu.Log.Debug("get task. [id: " + id.String() + "]")

	// タスクを取得
	task, err := tu.TaskRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Create はタスクを登録する
func (tu *TaskUsecase) Create(name task.Name, des task.Description) (*entity.Task, error) {
	// タスクを作成
	task := entity.NewTask()
	task.Update(name, des)

	// 永続化
	if err := tu.TaskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

// Update はタスクを更新する
func (tu *TaskUsecase) Update(id task.ID, name task.Name, des task.Description) (*entity.Task, error) {
	// タスクを取得
	task, err := tu.TaskRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	// タスク更新
	task.Update(name, des)

	// 永続化
	if err := tu.TaskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

// Complate はタスクを完了にする
func (tu *TaskUsecase) Complate(id task.ID) error {
	// タスクを取得
	task, err := tu.TaskRepo.ReadByID(id)
	if err != nil {
		return err
	}

	// ステータス更新
	if err = task.Complate(); err != nil {
		return err
	}

	// 永続化
	if err := tu.TaskRepo.Update(task); err != nil {
		return err
	}

	return nil
}
