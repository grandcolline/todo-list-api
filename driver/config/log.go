package config

import "github.com/kelseyhightower/envconfig"

// LogConf ログの設定
type LogConf struct {
	Type  string `default:"row"`
	Level string `default:"debug"` // ログレベル（dubug/info/error）
}

// Init ログ設定を環境変数から取得します
func (conf *LogConf) Init() {
	err := envconfig.Process("log", conf)
	if err != nil {
		panic("LogConf processing error: "+err.Error())
	}
}
