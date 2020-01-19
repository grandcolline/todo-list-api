package errfmt

import "github.com/grandcolline/todo-list-api/util/errors/code"

type ErrFmt struct {
	Code code.Code
	Str  string
}

// ConvErr 型のコンバート時に出されるエラー
var Conv = ErrFmt{
	Code: code.BadParams,
	Str:  "failed to convert %s because %s", // $1:変換しようとした物・$2:エラー理由
}
