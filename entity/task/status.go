package task

import "github.com/grandcolline/todo-list-api/util/errors"

/*
 Status タスクステータス
 タスクステータスはタスクの状態を表す値オブジェクト
   - DOING:    作業中
   - COMPLATE: 完了済み
 の2つの状態を持つ
*/
type Status struct{ value string }

// FIXME: ここvarなのでパッケージ外から変更可能だけど大丈夫？
var (
	// DOING 作業中
	DOING Status = Status{value: "doing"}
	// Complate 完了済み
	COMPLATE Status = Status{value: "complate"}
)

// DefaultStatus はタスクステータスのデフォルト値
func DefaultStatus() Status {
	return DOING
}

// NewStatus はタスクステータスを生成する
func NewStatus() Status {
	return DefaultStatus()
}

// ToStatus はStringをタスクステータスに変換する
func ToStatus(s string) (Status, error) {
	switch s {
	case DOING.String():
		return DOING, nil
	case COMPLATE.String():
		return COMPLATE, nil
	default:
		return Status{}, errors.NewConvErr("task.status", "invalit")
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
