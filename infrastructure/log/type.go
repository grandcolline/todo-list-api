package log

import "fmt"

// fmtType フォーマット形式 (row/json)
type fmtType string

const (
	rowType  fmtType = "row"
	jsonType fmtType = "json"
)

// convFmtType はstringをフォーマット形式に変換する.
func convFmtType(s string) (fmtType, error) {
	switch s {
	case "row", "ROW", "Row":
		return rowType, nil
	case "json", "JSON", "Json":
		return jsonType, nil
	default:
		return "", fmt.Errorf("fmtType covert error: %s", s)
	}
}
