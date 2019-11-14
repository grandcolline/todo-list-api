package task

import "fmt"

// Status タスクステータス
type Status string

const (
	// Doing 作業中
	Doing    Status = "doing"
	// Complate 完了済み
	Complate Status = "complate"
)

// NewStatus はタスクステータスを生成する
func NewStatus() Status {
	return Doing
}

// ToStatus はStringをタスクステータスに変換する
func ToStatus(s string) (Status, error) {
	switch s {
	case "Doing", "doing":
		return Doing, nil
	case "Complate", "complate":
		return Complate, nil
	default:
		return NewStatus(), fmt.Errorf("convert error")
	}
}

// String はタスクステータスをstringを返す
func (s Status) String() string {
	return string(s)
}
