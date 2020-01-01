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

// conf アプリケーション設定
var conf config.AppConf

// Serve はサーバの起動を行います
func Serve() {
	// 設定の読み込み
	conf.Init()

	// ListenPortの作成
	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		// FIXME: どうにかする
		// log.Fatalf("faild to listen port: %v", err)
		panic("faild to listen port: %v")
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	// DB接続の設定
	ctx := context.Background()
	cli, err := firestore.NewClient(ctx, "project-test")
	taskGateway := gateway.NewTaskGateway(cli, ctx)

	// ロガーファクトリの作成
	loggerFactory := func(id string) logger.Logger {
		return log.NewLog(id, "debug", "row", os.Stdout)
	}

	// タスクコントローラの作成
	taskController := controller.NewTaskController(taskGateway, loggerFactory)

	pb.RegisterTaskServiceServer(server, taskController)

	go func() {
		// FIXME: fmtでいいのか問題を検討(つーか、ログ出てないよねw
		// log.Printf("start grpc Server port: %s", conf.Port)
		fmt.Printf("start grpc Server port: %s", conf.Port)
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
