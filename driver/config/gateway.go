package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// GatewayConf Gatewayの設定
type GatewayConf struct {
	ProjectID string `default:"todo"` // FireStoreのプロジェクトID
}

// Init Gatewayの設定を環境変数から取得します
func (conf *GatewayConf) Init() {
	err := envconfig.Process("gateway", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
