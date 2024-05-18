package logger

type ILogger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
}

type Field struct {
	Key       string
	Type      uint8
	Integer   int64
	String    string
	Interface interface{}
}