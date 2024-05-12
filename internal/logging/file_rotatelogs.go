package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level, director string, maxSize, maxAge, maxBackups int, logInConsole bool) (zapcore.WriteSyncer, error) {
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s-%s.log", director, time.Now().Format("2006-01-02/15-04-05"), level), // 日志文件路径
		MaxSize:    maxSize,                                                                                // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups,                                                                             // 日志文件最多保存多少个备份
		MaxAge:     maxAge,                                                                                 // 文件最多保存多少天
		Compress:   true,                                                                                   // 是否压缩
	}

	var err error
	// fileWriter, err := rotatelogs.New(
	// 	path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
	// 	rotatelogs.WithClock(rotatelogs.Local),
	// 	rotatelogs.WithMaxAge(time.Duration(global.CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
	// 	rotatelogs.WithRotationTime(time.Hour*24),
	// )
	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), err
	}
	return zapcore.AddSync(&hook), err
}
