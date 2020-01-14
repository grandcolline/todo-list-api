package log

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/grandcolline/todo-list-api/usecase/logger"
)

// Log ロガー実装
type Log struct {
	tranID  string
	level   level
	fmtType fmtType
	writer  io.Writer
}

// NewLog ロガー実装を作成する.
func NewLog(id, levelStr, fmtTypeStr string, writer io.Writer) logger.Logger {
	// FIXME: トランザクションIDが空の場合のハンドリング

	// 出力レベルを取得
	level, err := convLv(levelStr)
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
		level:   level,
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
	if lv < l.level {
		return
	}

	// 整形して出力
	switch l.fmtType {
	case row:
		fmt.Fprintln(l.writer, fmt.Sprintf(
			"%s [%s][%s] %s",
			time.Now().Format("2006/01/02 15:04:05"),
			lv.string(),
			l.tranID,
			s,
		))
	case stackdriver:
		entry := map[string]string{
			"time":     time.Now().Format(time.RFC3339Nano),
			"severity": lv.string(),
			"message":  s,
			"tranID":   l.tranID,
		}
		b, _ := json.Marshal(entry)
		fmt.Fprintln(l.writer, string(b))
	}
}
