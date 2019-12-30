package gateway

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/infrastructure/repository/gateway/collection"
	"github.com/grandcolline/todo-list-api/usecase/repository"
	"github.com/grandcolline/todo-list-api/util/errors"
)

// TaskGateway タスクレポジトリ実装
type TaskGateway struct {
	cli *firestore.Client
	ctx context.Context
}

// NewTaskGateway タスクレポジトリ実装を作成する
func NewTaskGateway(cli *firestore.Client, ctx context.Context) repository.TaskRepository {
	return &TaskGateway{
		cli: cli,
		ctx: ctx,
	}
}

// Upsert タスクを登録/更新をする
func (tg *TaskGateway) Upsert(task *entity.Task) error {
	// タスクエンティティをコレクションに詰め替え
	var tc collection.TaskCollection
	id := tc.FromEntity(task)

	// FireStoreに保存
	if _, err := tg.cli.Collection(tc.CollectionName()).Doc(id).Set(tg.ctx, tc); err != nil {
		return errors.Errorf(errors.Database, "failed to upsert task: %s", err)
	}

	return nil
}

// ReadByID タスクをIDで取得する
func (tg *TaskGateway) ReadByID(taskID task.ID) (*entity.Task, error) {
	var tc collection.TaskCollection
	snapshot, err := tg.cli.Collection(tc.CollectionName()).Doc(taskID.String()).Get(tg.ctx)
	if err != nil {
		return nil, errors.Errorf(errors.Database, "failed to read by id task: %s", err)
	}

	if err = snapshot.DataTo(&tc); err != nil {
		return nil, errors.Errorf(errors.Database, "failed to read by id task: %s", err)
	}
	return tc.ToEntity(taskID.String())
}

// Delete タスクを削除する
func (tg *TaskGateway) Delete(taskID task.ID) error {
	var tc collection.TaskCollection
	_, err := tg.cli.Collection(tc.CollectionName()).Doc(taskID.String()).Delete(tg.ctx)
	if err != nil {
		return err
	}
	return nil
}
