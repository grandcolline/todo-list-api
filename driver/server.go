package driver

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/grandcolline/todo-list-api/adapter/controller"
	"github.com/grandcolline/todo-list-api/adapter/controller/proto/pb"
	"github.com/grandcolline/todo-list-api/driver/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

var (
	taskController *controller.TaskController
)

func InitController() {
	taskController = controller.NewTaskController(taskRepo, loggerFactory)
}

// Serve はサーバの起動を行います
func Serve() {
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
