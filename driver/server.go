package driver

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/grandcolline/todo-list-api/application/controller"
	"github.com/grandcolline/todo-list-api/application/controller/proto/pb"
	"github.com/grandcolline/todo-list-api/driver/config"
	"github.com/grandcolline/todo-list-api/infrastructure/repository/gateway"
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
		log.Fatalf("faild to listen port: %v", err)
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
	// gatewayNewTaskRepoImpl

	// ロガーの作成
	// logger := logger.NewLogger(conf.LogLevel)

	// タスクコントローラの作成
	taskController := controller.NewTaskController(taskGateway)

	pb.RegisterTaskServiceServer(server, taskController)

	go func() {
		log.Printf("start grpc Server port: %s", conf.Port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc Server...")
	server.GracefulStop()

}
