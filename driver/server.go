package driver

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/grandcolline/todo-list-api/application/controller"
	"github.com/grandcolline/todo-list-api/application/controller/proto/pb"
	"github.com/grandcolline/todo-list-api/driver/config"
	"github.com/grandcolline/todo-list-api/infrastructure/gateway"
	"github.com/grandcolline/todo-list-api/infrastructure/log"
	"github.com/grandcolline/todo-list-api/usecase/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	"cloud.google.com/go/firestore"
)

// Serve はサーバの起動を行います
func Serve() {

	// ロガーファクトリの作成
	var logConf config.LogConf
	logConf.Init()
	loggerFactory := func(id string) logger.Logger {
		return log.NewLog(id, logConf.Level, logConf.Type, os.Stdout)
	}

	// DB接続の設定
	var firestoreConf config.FirestoreConf
	firestoreConf.Init()
	ctx := context.Background()
	cli, err := firestore.NewClient(ctx, firestoreConf.ProjectID)
	taskGateway := gateway.NewTaskGateway(cli, ctx)

	// タスクコントローラの作成
	taskController := controller.NewTaskController(taskGateway, loggerFactory)

	// アプリケーション設定の読み込み
	var appConf config.AppConf
	appConf.Init()

	// ListenPortの作成
	lis, err := net.Listen("tcp", ":"+appConf.Port)
	if err != nil {
		panic("faild to listen port: " + err.Error())
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	pb.RegisterTaskServiceServer(server, taskController)

	go func() {
		fmt.Printf("start grpc Server port: %s\n", appConf.Port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	// FIXME: fmtでいいのか問題を検討(つーか、ログ出てないよねw
	// log.Println("stopping grpc Server...")
	fmt.Println("stopping grpc Server...")
	server.GracefulStop()

}
