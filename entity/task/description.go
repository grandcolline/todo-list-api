package task

import (
	"unicode/utf8"

	"github.com/grandcolline/todo-list-api/util/errors"
)

/*
 Description タスク詳細
 タスク詳細は300文字以内の文字列の値オブジェクト
*/
type Description struct{ value string }

// NewDescription はタスク詳細を生成する
func NewDescription() Description {
	return Description{value: ""}
}

// ToDescription はStringをタスク詳細に変換する
func ToDescription(s string) (Description, error) {
	// 文字数チェック
	if utf8.RuneCountInString(s) > 300 {
		return Description{}, errors.New(errors.BadParams, "failed to convert task.description: over 300")
	}

	return Description{value: s}, nil
}

// String はタスク詳細をstringで返す
func (d *Description) String() string {
	return d.value
}
