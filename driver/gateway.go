package driver

import (
	"context"

	"github.com/grandcolline/todo-list-api/adapter/gateway"
	"github.com/grandcolline/todo-list-api/driver/config"
	"github.com/grandcolline/todo-list-api/usecase/repository"

	"cloud.google.com/go/firestore"
)

var (
	taskRepo repository.Task
)

func InitGateway() {
	// 環境変数の取得
	var firestoreConf config.FirestoreConf
	firestoreConf.Init()

	// DB接続
	ctx := context.Background()
	cli, err := firestore.NewClient(ctx, firestoreConf.ProjectID)
	if err != nil {
		panic(err)
	}

	// Gatewayの作成
	taskRepo = gateway.NewTask(cli, ctx)
}
