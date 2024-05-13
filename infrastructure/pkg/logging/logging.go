package logging

import (
	"time"

	"github.com/gookit/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILoggerCofig interface {
	LevelParam() zapcore.Level
	PrefixParam() string
	FormatParam() string
	DirectorParam() string
	EncodeLevelParam() zapcore.LevelEncoder
	StacktraceKeyParam() string
	MaxAgeParam() int
	MaxSizeParam() int
	MaxBackupsParam() int
	LogInConsoleParam() bool
}

func NewLoggingEnforce(conf ILoggerCofig) *zap.Logger {
	ZapImpl := _zapImpl{conf}
	cores := ZapImpl.GetZapCores()
	return zap.New(zapcore.NewTee(cores...))
}

// var Zap = new(_zapImpl)

type _zapImpl struct {
	ILoggerCofig
}

func (z *_zapImpl) GetEncoder() zapcore.Encoder {
	if z.FormatParam() == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

func (z *_zapImpl) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  z.StacktraceKeyParam(),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.EncodeLevelParam(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func (z *_zapImpl) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotatelogs.GetWriteSyncer(l.String(), z.DirectorParam(), z.MaxSizeParam(), z.MaxAgeParam(), z.MaxBackupsParam(), z.LogInConsoleParam()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		color.Error.Printf("Get Write Syncer Failed err:%v\n", err.Error())
		return nil
	}

	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zapImpl) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(z.PrefixParam() + "2006/01/02 - 15:04:05.000"))
}

func (z *_zapImpl) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.LevelParam(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

func (z *_zapImpl) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
