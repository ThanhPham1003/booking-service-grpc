package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger zap.Logger
}

func NewLogger()ILogger {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		return nil
	}
	return &ZapLogger{logger: *logger}
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	// mapFields := make(map[string]interface{})
	// mapFields := fieldsToMap(fields)
	l.logger.Debug(msg)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg)
}
func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg)
}

// func fieldsToMap(fields ...Field) map[string]interface{} {
// 	result := make(map[string]interface{}, len(fields))
// 	for _, field := range fields {
// 		result[field.Key] = field.Value
// 	}
// 	return result
// }
