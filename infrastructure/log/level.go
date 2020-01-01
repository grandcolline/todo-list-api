package log

import "fmt"

// level ログレベル
type level uint

const (
	// iotaでログの優先度が低い順に並べる.
	debugLv level = iota + 1
	infoLv
	errorLv
)

// convLv はstringをログレベルに変換する.
func convLv(s string) (level, error) {
	switch s {
	case "debug", "DEBUG", "Debug":
		return debugLv, nil
	case "info", "INFO", "Info":
		return infoLv, nil
	case "error", "ERROR", "Error":
		return errorLv, nil
	default:
		return 0, fmt.Errorf("level covert error: %s", s)
	}
}

func (lv level) string() string {
	switch lv {
	case debugLv:
		return "DEBUG"
	case infoLv:
		return "INFO"
	case errorLv:
		return "ERROR"
	default:
		return ""
	}
}

func (lv level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + lv.string() + `"`), nil
}
