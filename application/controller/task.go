package controller

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"github.com/grandcolline/todo-list-api/application/controller/proto/pb"
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/usecase"
	"github.com/grandcolline/todo-list-api/usecase/repository"
)

// TaskController タスクコントローラ
type TaskController struct {
	taskUc *usecase.TaskUsecase
}

// NewTaskController はタスクコントローラを作成する
func NewTaskController(repo repository.TaskRepository) *TaskController {
	return &TaskController{
		taskUc: usecase.NewTaskUsecase(repo),
	}
}

// GetTask タスク取得
func (tc *TaskController) GetTask(c context.Context, p *pb.GetTaskRequest) (*pb.Task, error) {
	id, err := task.ToID(p.ID)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	taskEnt, err := tc.taskUc.GetByID(id)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	return convTask(taskEnt), nil
}

// CreateTask タスク登録
func (tc *TaskController) CreateTask(c context.Context, p *pb.CreateTaskRequest) (*pb.Task, error) {
	name, err := task.ToName(p.Name)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	des, err := task.ToDescription(p.Description)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	taskEnt, err := tc.taskUc.Create(name, des)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	return convTask(taskEnt), nil
}

// UpdateTask タスク更新
func (tc *TaskController) UpdateTask(c context.Context, p *pb.UpdateTaskRequest) (*empty.Empty, error) {
	id, err := task.ToID(p.ID)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	name, err := task.ToName(p.Name)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	des, err := task.ToDescription(p.Description)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	if err := tc.taskUc.Update(id, name, des); err != nil {
		// FIXME: error handling
		return nil, err
	}
	return &empty.Empty{}, nil
}

// CompleteTask タスク完了
func (tc *TaskController) CompleteTask(c context.Context, p *pb.CompleteTaskRequest) (*empty.Empty, error) {
	id, err := task.ToID(p.ID)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	if err := tc.taskUc.Complate(id); err != nil {
		// FIXME: error handling
		return nil, err
	}
	return &empty.Empty{}, nil
}

// タスクエンティティをタスク共通メッセージに変換します
func convTask(e *entity.Task) *pb.Task {
	return &pb.Task{
		ID:          e.ID.String(),
		Name:        e.Name.String(),
		Description: e.Description.String(),
		Status:      convStatus(e.Status),
	}
}

// ステータスをenumに変換
func convStatus(s task.Status) pb.Status {
	switch s {
	case task.Complate:
		return pb.Status_COMPLETE
	default:
		return pb.Status_DOING
	}
}
