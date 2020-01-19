package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// LogConf ログの設定
type LogConf struct {
	Type  string `default:"row"`
	Level string `default:"debug"` // ログレベル（dubug/info/error）
}

// Init アプリケーション全体設定を環境変数から取得します
func (conf *LogConf) Init() {
	err := envconfig.Process("log", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
