package task

import (
	"fmt"
	"unicode/utf8"
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
		// FIXME: エラーの共通化
		return NewDescription(), fmt.Errorf("convert error")
	}

	return Description{value: s}, nil
}

// String はタスク詳細をstringで返す
func (d Description) String() string {
	return d.value
}
