package task

import (
	"unicode/utf8"

	"github.com/grandcolline/todo-list-api/util/errors"
	"github.com/grandcolline/todo-list-api/util/errors/errfmt"
)

// Name タスク名
type Name struct{ value string }

const (
	// MinNameLen タスク名の最小文字数
	MinNameLen int = 1
	// MaxNameLen タスク名の最大文字数
	MaxNameLen int = 50
)

// DefaultName はタスク名のデフォルト値を返す
func DefaultName() Name {
	return Name{value: "New Task"}
}

// NewName はタスク名を生成する
func NewName() Name {
	return DefaultName()
}

// ToName はStringをタスク名に変換する
func ToName(s string) (Name, error) {
	// 文字数チェック
	if utf8.RuneCountInString(s) < MinNameLen {
		return Name{}, errors.NewFromFmt(errfmt.Conv, "task.name", "less min length")
	}
	if MaxNameLen < utf8.RuneCountInString(s) {
		return Name{}, errors.NewFromFmt(errfmt.Conv, "task.name", "over max length")
	}

	return Name{value: s}, nil
}

// String はタスク名をstringを返す
func (n Name) String() string {
	return n.value
}
