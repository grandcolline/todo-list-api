package gateway

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/grandcolline/todo-list-api/entity"
	"github.com/grandcolline/todo-list-api/entity/task"
	"github.com/grandcolline/todo-list-api/infrastructure/gateway/collection"
	"github.com/grandcolline/todo-list-api/usecase/repository"
	"github.com/grandcolline/todo-list-api/util/errors"
	"github.com/grandcolline/todo-list-api/util/errors/code"
)

// Task タスクレポジトリ実装
type Task struct {
	cli *firestore.Client
	ctx context.Context
}

// NewTask タスクレポジトリ実装を作成する
func NewTask(cli *firestore.Client, ctx context.Context) repository.Task {
	return &Task{
		cli: cli,
		ctx: ctx,
	}
}

// Upsert タスクを登録/更新をする
func (t *Task) Upsert(task *entity.Task) error {
	// タスクエンティティをコレクションに詰め替え
	var tc collection.TaskCollection
	id := tc.FromEntity(task)

	// FireStoreに保存
	if _, err := t.cli.Collection(tc.CollectionName()).Doc(id).Set(t.ctx, tc); err != nil {
		return errors.New(code.Database, "failed to upsert task("+id+"): "+err.Error())
	}

	return nil
}

// ReadByID タスクをIDで取得する
func (t *Task) ReadByID(taskID task.ID) (*entity.Task, error) {
	var tc collection.TaskCollection
	snapshot, err := t.cli.Collection(tc.CollectionName()).Doc(taskID.String()).Get(t.ctx)
	if err != nil {
		// FIXME notFoundのハンドリング
		return nil, errors.New(code.Database, "failed to read by id task("+taskID.String()+"): "+err.Error())
	}

	if err = snapshot.DataTo(&tc); err != nil {
		// FIXME: 上と同じエラーじゃダメやない？
		return nil, errors.New(code.Database, "failed to read by id task("+taskID.String()+"): "+err.Error())
	}
	return tc.ToEntity(taskID.String())
}

// Delete タスクを削除する
func (t *Task) Delete(taskID task.ID) error {
	var tc collection.TaskCollection
	_, err := t.cli.Collection(tc.CollectionName()).Doc(taskID.String()).Delete(t.ctx)
	if err != nil {
		return errors.New(code.Database, "failed to delete task: "+err.Error())
	}
	return nil
}
