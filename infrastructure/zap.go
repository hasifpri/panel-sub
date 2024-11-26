package infrastructure

import (
	"fmt"
	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
	"strings"
)

var ZAPLOGGER *zap.Logger

func InitializeZap() {

	var host string
	var err error
	var coreAll zapcore.Core

	if host, err = os.Hostname(); err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	defaultLogLevel := zapcore.DebugLevel
	coreAll = zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel)

	if strings.Contains(infrastructureconfiguration.LoggingTool, "STDOUT") && strings.Contains(infrastructureconfiguration.LoggingTool, "GRAYLOG") {
		coreGrayLog, err := gelf.NewCore(
			gelf.Addr(infrastructureconfiguration.GelfAddr),
			gelf.Host(host),
		)
		if err != nil {
			panic(err)
		} else {
			coreAll = zapcore.NewTee(
				zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), coreAll),
				coreGrayLog,
			)
		}
	}

	var zapLogger = zap.New(
		coreAll,
		zap.AddCaller(),
		zap.AddStacktrace(zap.LevelEnablerFunc(func(level zapcore.Level) bool { return coreAll.Enabled(1) })),
	)

	ZAPLOGGER = zapLogger

	if err != nil {
		fmt.Println("=== Load Logger Zap ===")
		fmt.Println("error InitializeZap() : ", err)
		fmt.Println("=== End Load Logger Zap ===")
	} else {
		fmt.Println("=== Load Logger Zap ===")
		fmt.Printf("Zap Logger Is Ready : %+v\n", ZAPLOGGER)
		fmt.Println("=== End Load Logger Zap ===")
	}
}
