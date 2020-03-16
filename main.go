package main

import (
	"github.com/grandcolline/todo-list-api/driver"
)

func main() {
	// ロガーファクトリの作成
	driver.InitLog()
	// Gatewayの作成
	driver.InitGateway()
	// タスクコントローラの作成
	driver.InitController()

	// サーバの起動
	driver.Serve()
}
