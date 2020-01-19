package errors

import "github.com/grandcolline/todo-list-api/util/errors/code"

type errMsg struct {
	Code code.Code
	Msg  string
}

// Conv 型のコンバート時に出されるエラー
var Conv = errMsg{
	Code: code.BadParams,
	Msg:  "failed to convert %s because %s", // $1:変換しようとした物・$2:エラー理由
}
