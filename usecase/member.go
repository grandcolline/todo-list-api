package usecase

import (
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"

	// "github.com/grandcolline/todo-list-api/usecase/logger"
	"github.com/grandcolline/todo-list-api/usecase/repository"
)

// MemberUsecase メンバーユースケース
type MemberUsecase struct {
	taskRepo repository.TaskRepository
	// Log      logger.Logger
}

// NewMemberUsecase はメンバーユースケースを作成する
func NewMemberUsecase(taskRepo repository.TaskRepository) *MemberUsecase {
	return &MemberUsecase{
		taskRepo: taskRepo,
	}
}

// GetByID はIDでタスクを取得する
func (mu *MemberUsecase) GetByID(id task.ID) (*entity.Task, error) {

	// タスクを取得
	task, err := mu.taskRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Create はタスクを登録する
func (mu *MemberUsecase) Create(name task.Name, des task.Description) (*entity.Task, error) {
	// タスクを作成
	task := entity.NewTask()
	task.Update(name, des)

	// 永続化
	if err := mu.taskRepo.Upsert(task); err != nil {
		return nil, err
	}

	return task, nil
}

// Update はタスクを更新する
func (mu *MemberUsecase) Update(id task.ID, name task.Name, des task.Description) error {
	// タスクを取得
	task, err := mu.taskRepo.ReadByID(id)
	if err != nil {
		return err
	}

	// タスク更新
	task.Update(name, des)

	// 永続化
	if err := mu.taskRepo.Upsert(task); err != nil {
		return err
	}

	return nil
}

// Complate はタスクを完了にする
func (mu *MemberUsecase) Complate(id task.ID) error {
	// タスクを取得
	task, err := mu.taskRepo.ReadByID(id)
	if err != nil {
		return err
	}

	// ステータス更新
	if err = task.Complate(); err != nil {
		return err
	}

	// 永続化
	if err := mu.taskRepo.Upsert(task); err != nil {
		return err
	}

	return nil
}
