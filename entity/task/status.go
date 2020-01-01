package task

import "github.com/grandcolline/todo-list-api/util/errors"

/*
 Status タスクステータス
 タスクステータスはタスクの状態を表す値オブジェクト
   - Doing 作業中
   - Complate 完了済み
 の2つの状態を持つ
*/
type Status struct{ value string }

// Enum
var (
	// DOING 作業中
	DOING Status = Status{value: "doing"}
	// Complate 完了済み
	COMPLATE Status = Status{value: "complate"}
)

// NewStatus はタスクステータスを生成する
func NewStatus() Status {
	return DOING
}

// ToStatus はStringをタスクステータスに変換する
func ToStatus(s string) (Status, error) {
	switch s {
	case "Doing", "doing", "DOING":
		return DOING, nil
	case "Complate", "complate", "COMPLATE":
		return COMPLATE, nil
	default:
		return NewStatus(), errors.New(errors.BadParams, "failed to convert task.status: invalit")
	}
}

// IsDoing は作業中かどうかを返す
func (s Status) IsDoing() bool {
	return s.value == DOING.value
}

// IsComplate は完了かどうかを返す
func (s Status) IsComplate() bool {
	return s.value == COMPLATE.value
}

// String はタスクステータスをstringで返す
func (s Status) String() string {
	return s.value
}
