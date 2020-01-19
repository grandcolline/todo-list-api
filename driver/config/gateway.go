package config

import (
	// "log"

	"github.com/kelseyhightower/envconfig"
)

// GatewayConf Gatewayの設定
type FirestoreConf struct {
	Host      string `default:"firestore"` // ホスト
	Port      string `default:"8081"`      // ポート
	ProjectID string `default:"todo"`      // プロジェクトID
	Emulator  bool   `default:"false"`     // エミュレータかどうか
}

// Init Gatewayの設定を環境変数から取得します
func (conf *FirestoreConf) Init() {
	err := envconfig.Process("firestore", conf)
	if err != nil {
		panic("GatewayConf processing error: "+err.Error())
	}
}
