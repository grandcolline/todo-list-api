package log

import "fmt"

// fmtType フォーマット形式
type fmtType string

const (
	row         fmtType = "row"
	stackdriver fmtType = "json"
)

// convFmtType はstringをフォーマット形式に変換する.
func convFmtType(s string) (fmtType, error) {
	switch s {
	case "row":
		return row, nil
	case "stackdriver":
		return stackdriver, nil
	default:
		return "", fmt.Errorf("fmtType covert error: %s", s)
	}
}
