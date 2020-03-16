package driver

import (
	"os"

	"github.com/grandcolline/todo-list-api/driver/config"
	"github.com/grandcolline/todo-list-api/infrastructure/log"
	"github.com/grandcolline/todo-list-api/usecase/logger"
)

var loggerFactory func(id string) logger.Logger

func InitLog(){
	// 環境変数の取得
	var logConf config.LogConf
	logConf.Init()

	// ロガーファクトリーの作成
	loggerFactory = func(id string) logger.Logger {
		return log.NewLog(id, logConf.Level, logConf.Type, os.Stdout)
	}
}
