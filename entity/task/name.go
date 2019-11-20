package task

import (
	"fmt"
	"unicode/utf8"
)

/*
 Name タスク名
 タスク名は1文字以上50文字以内の文字列の値オブジェクト
*/
type Name string

// NewName はタスク名を生成する
func NewName() Name {
	return Name("New Task")
}

// ToName はStringをタスク名に変換する.
func ToName(s string) (Name, error) {
	// 空文字チェック
	if s == "" {
		// FIXME: エラーの共通化
		return NewName(), fmt.Errorf("convert error")
	}
	// 文字数チェック
	if utf8.RuneCountInString(s) > 50 {
		// FIXME: エラーの共通化
		return NewName(), fmt.Errorf("convert error")
	}

	return Name(s), nil
}

// String はタスク名をstringを返す
func (n Name) String() string {
	return string(n)
}
