package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
)

type LogConfig struct {
	LogLevel          string // 日志等级 debug | info | warn | error
	LogFormat         string // 日志格式 logfmt | json
	LogPath           string // 日志文件路径
	LogName           string // 日志文件名称
	LogFileMaxSize    int    // 【日志分割】单个日志大小（MB）
	LogFileMaxBackups int    // 【日志分割】备份日志保留数量
	LogFileMaxAge     int    // 【日志分割】备份日志保留时间（天）
	LogCompress       bool   // 是否压缩日志
	LogStdout         bool   // 是否同时输出到控制台
}

func InitAppLogger(conf *LogConfig) {
	// ======== writer ========
	var writer zapcore.WriteSyncer
	// 创建日志路径
	if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
		log.Panic("init app logger error:", err)
	}
	// 日志分割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogName),
		MaxSize:    conf.LogFileMaxSize,
		MaxBackups: conf.LogFileMaxBackups,
		MaxAge:     conf.LogFileMaxAge,
		Compress:   conf.LogCompress,
	}
	if conf.LogStdout {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		writer = zapcore.AddSync(lumberJackLogger)
	}

	// ======== level ========
	var level zapcore.Level
	switch conf.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// ======== encode ========
	var encoder zapcore.Encoder
	encoding := zap.NewProductionEncoderConfig()
	encoding.EncodeTime = zapcore.ISO8601TimeEncoder
	encoding.EncodeLevel = zapcore.CapitalLevelEncoder
	if conf.LogFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encoding)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoding)
	}

	// ======== logger ========
	logger := zap.New(zapcore.NewCore(encoder, writer, level))
	zap.ReplaceGlobals(logger)
}

func DEBUG_MSG(args ...any) {
	zap.S().Debug(args...)
}

func DEBUG_MSG_F(template string, args ...any) {
	zap.S().Debugf(template, args...)
}

func INFO_MSG(args ...any) {
	zap.S().Info(args...)
}

func INFO_MSG_F(template string, args ...any) {
	zap.S().Infof(template, args...)
}

func WARN_MSG(args ...any) {
	zap.S().Warn(args...)
}

func WARN_MSG_F(template string, args ...any) {
	zap.S().Warnf(template, args...)
}

func ERROR_MSG(args ...any) {
	zap.S().Error(args...)
}

func ERROR_MSG_F(template string, args ...any) {
	zap.S().Errorf(template, args...)
}

func PANIC_MSG(args ...any) {
	zap.S().Panic(args...)
}

func PANIC_MSG_F(template string, args ...any) {
	zap.S().Panicf(template, args...)
}
