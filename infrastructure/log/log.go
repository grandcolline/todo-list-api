package log

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"time"

	"github.com/grandcolline/todo-list-api/usecase/logger"
)

// Log ロガー実装
type Log struct {
	tranID  string
	outLv   level
	fmtType fmtType
	writer  io.Writer
}

// logMessage ログ表示内容
type logMessage struct {
	tranID    string
	filename  string
	line      int
	message   string
	level     level
	timestamp time.Time
}

// NewLog ロガー実装を作成する.
func NewLog(id, outLvStr, fmtTypeStr string, writer io.Writer) logger.Logger {

	// TODO: トランザクションIDが空の場合のハンドリング

	// 出力レベルを取得
	outLv, err := convLv(outLvStr)
	if err != nil {
		panic(err)
	}
	// 出力形式を取得
	fmtType, err := convFmtType(fmtTypeStr)
	if err != nil {
		panic(err)
	}
	return &Log{
		tranID:  id,
		outLv:   outLv,
		fmtType: fmtType,
		writer:  writer,
	}
}

// Debug はDEBUGレベルのログを出力する
func (l *Log) Debug(s string) {
	l.output(s, debugLv)
}

// Info はINFOレベルのログを出力する
func (l *Log) Info(s string) {
	l.output(s, infoLv)
}

// Error はエラーレベルのログを出力する
func (l *Log) Error(s string) {
	l.output(s, errorLv)
}

// fakeTime テスト用時刻固定時間
var fakeTime time.Time

// output はログを整形して出力する
func (l *Log) output(s string, lv level) {
	// アウトプットレベルより小さい時は出力せずに終了
	if lv < l.outLv {
		return
	}

	_, filename, line, _ := runtime.Caller(2)

	// コンテンツを作成
	lm := logMessage{
		tranID:    l.tranID,
		timestamp: time.Now(),
		level:     lv,
		message:   s,
		filename:  filename,
		line:      line,
	}

	// format
	var res string
	switch l.fmtType {
	case rowType:
		res = fmt.Sprintf(
			"%s [%s][%s](%s:%s) %s",
			lm.timestamp.Format("2006/01/02 15:04:05"),
			lm.level.string(),
			lm.tranID,
			lm.filename,
			strconv.Itoa(lm.line),
			lm.message,
		)
	case jsonType:
		b, _ := json.Marshal(lm)
		res = string(b)
	}

	// output
	fmt.Fprintln(l.writer, res)
}
