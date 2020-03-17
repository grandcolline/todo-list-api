package controller

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"github.com/grandcolline/todo-list-api/application/controller/proto/pb"
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/usecase"
	"github.com/grandcolline/todo-list-api/usecase/logger"
	"github.com/grandcolline/todo-list-api/usecase/repository"
	"github.com/grandcolline/todo-list-api/util/errors"
)

// TaskController タスクコントローラ
type TaskController struct {
	userUc        *usecase.User
	loggerFactory func(string) logger.Logger
}

// NewTaskController はタスクコントローラを作成する
func NewTaskController(taskRepo repository.Task, loggerFactory func(string) logger.Logger) *TaskController {
	return &TaskController{
		userUc:        usecase.NewUser(taskRepo),
		loggerFactory: loggerFactory,
	}
}

// GetTask タスク取得
func (tc *TaskController) GetTask(c context.Context, p *pb.GetTaskRequest) (*pb.Task, error) {
	// loggerの作成
	log := tc.loggerFactory("") // FIXME: IDをしっかり入れる

	id, err := task.ToID(p.ID)
	if err != nil {
		log.Error(errors.Format(err))
		return nil, err
	}
	taskEnt, err := tc.userUc.GetByID(id)
	if err != nil {
		log.Error(errors.Format(err))
		return nil, err
	}
	return convTask(taskEnt), nil
}

// CreateTask タスク登録
func (tc *TaskController) CreateTask(c context.Context, p *pb.CreateTaskRequest) (*pb.Task, error) {
	// loggerの作成
	log := tc.loggerFactory("") // FIXME: IDをしっかり入れる

	// 型変換
	name, err := task.ToName(p.Name)
	if err != nil {
		log.Error(errors.Format(err))
		return nil, err
	}
	des, err := task.ToDescription(p.Description)
	if err != nil {
		log.Error(errors.Format(err))
		return nil, err
	}

	// Usecaseの呼び出し
	taskEnt, err := tc.userUc.Create(name, des)
	if err != nil {
		return nil, err
	}

	// response
	log.Info("task created. id: " + taskEnt.ID.String())
	return convTask(taskEnt), nil
}

// UpdateTask タスク更新
func (tc *TaskController) UpdateTask(c context.Context, p *pb.UpdateTaskRequest) (*empty.Empty, error) {
	// loggerの作成
	// log := tc.loggerFactory("") // FIXME: IDをしっかり入れる

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
	if err := tc.userUc.Update(id, name, des); err != nil {
		// FIXME: error handling
		return nil, err
	}
	return &empty.Empty{}, nil
}

// CompleteTask タスク完了
func (tc *TaskController) CompleteTask(c context.Context, p *pb.CompleteTaskRequest) (*empty.Empty, error) {
	// loggerの作成
	// log := tc.loggerFactory("") // FIXME: IDをしっかり入れる

	id, err := task.ToID(p.ID)
	if err != nil {
		// FIXME: error handling
		return nil, err
	}
	if err := tc.userUc.Complate(id); err != nil {
		// FIXME: error handling
		return nil, err
	}
	return &empty.Empty{}, nil
}

// convTask タスクエンティティをタスク共通メッセージに変換します
func convTask(e *entity.Task) *pb.Task {
	return &pb.Task{
		ID:          e.ID.String(),
		Name:        e.Name.String(),
		Description: e.Description.String(),
		Status:      convStatus(e.Status),
	}
}

// convStatus ステータスをenumに変換
func convStatus(s task.Status) pb.Status {
	switch {
	case s.IsComplate():
		return pb.Status_COMPLETE
	default:
		return pb.Status_DOING
	}
}
