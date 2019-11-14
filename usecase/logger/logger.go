package logger

// Logger ロガー
type Logger interface {
	Debug(string)
	Info(string)
	Error(string)
}
