package task

import (
	"unicode/utf8"

	"github.com/grandcolline/todo-list-api/util/errors"
)

/*
 Name タスク名
 タスク名は1文字以上50文字以内の文字列の値オブジェクト
*/
type Name struct{ value string }

// NewName はタスク名を生成する
func NewName() Name {
	return Name{value: "New Task"}
}

// ToName はStringをタスク名に変換する
func ToName(s string) (Name, error) {
	// 空文字チェック
	if s == "" {
		return NewName(), errors.New(errors.BadParams, "failed to convert task.name: empty")
	}
	// 文字数チェック
	if utf8.RuneCountInString(s) > 50 {
		return NewName(), errors.New(errors.BadParams, "failed to convert task.name: over 50")
	}

	return Name{value: s}, nil
}

// String はタスク名をstringを返す
func (n Name) String() string {
	return n.value
}
