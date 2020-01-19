package task

import (
	"unicode/utf8"

	"github.com/grandcolline/todo-list-api/util/errors"
)

// Description タスク詳細
type Description struct{ value string }

// DefaultDescription タスク詳細のデフォルト値
func DefaultDescription() Description {
	return Description{value: ""}
}

const (
	// MinDescriptionLen タスク詳細の最小文字数
	MinDescriptionLen int = 0
	// MaxDescriptionLen タスク詳細の最大文字数
	MaxDescriptionLen int = 300
)

// NewDescription はタスク詳細を生成する
func NewDescription() Description {
	return DefaultDescription()
}

// ToDescription はStringをタスク詳細に変換する
func ToDescription(s string) (Description, error) {
	// 文字数チェック
	if utf8.RuneCountInString(s) > MaxDescriptionLen {
		return Description{}, errors.NewFromFmt(errors.Conv, "task.description", "over max length")
	}

	return Description{value: s}, nil
}

// String はタスク詳細をstringで返す
func (d *Description) String() string {
	return d.value
}
