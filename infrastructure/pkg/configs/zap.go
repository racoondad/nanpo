package configs

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level         string `mapstructure:"level" yaml:"level"`                   // 级别
	Prefix        string `mapstructure:"prefix" yaml:"prefix"`                 // 日志前缀
	Format        string `mapstructure:"format" yaml:"format"`                 // 输出
	Director      string `mapstructure:"director" yaml:"director"`             // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" yaml:"encode-level"`     // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	MaxAge        int    `mapstructure:"max-age" yaml:"max-age"`               // 日志留存时间
	MaxSize       int    `mapstructure:"max-size" yaml:"max-size"`             // 最大尺寸
	MaxBackups    int    `mapstructure:"max-backups" yaml:"max-backups"`       // 最大被分数
	LogInConsole  bool   `mapstructure:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

func (z *Zap) EncodeLevelParam() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (z *Zap) LevelParam() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (z *Zap) PrefixParam() string {
	return z.Prefix
}

func (z *Zap) FormatParam() string {
	return z.Format
}

func (z *Zap) DirectorParam() string {
	return z.Director
}

func (z *Zap) StacktraceKeyParam() string {
	return z.StacktraceKey
}

func (z *Zap) MaxAgeParam() int {
	return z.MaxAge
}

func (z *Zap) MaxSizeParam() int {
	return z.MaxSize
}

func (z *Zap) MaxBackupsParam() int {
	return z.MaxBackups
}

func (z *Zap) LogInConsoleParam() bool {
	return z.LogInConsole
}
