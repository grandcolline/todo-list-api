package config

import "github.com/kelseyhightower/envconfig"

// AppConf アプリケーション全体の設定
type AppConf struct {
	Port string `default:"50051"` // サーバ起動時に受け付けるポート
	// LogLevel string `default:"INFO"`  // ログレベル（DEBUG/INFO/ERROR）
}

// Init アプリケーション全体設定を環境変数から取得します
func (conf *AppConf) Init() {
	err := envconfig.Process("app", conf)
	if err != nil {
		panic("AppConf processing error: " + err.Error())
	}
}
