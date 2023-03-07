import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type openapiartlog struct {
	RootDir              *string
	LogDir               *string
	Debug                *bool
	DisableStdOutLogging *bool
	MaxLogSizeMB         *int
	MaxLogBackups        *int
}

var (
	Openapiartlog  *openapiartlog
	GlobalLogger   *zerolog.Logger
	GlobalLogLevel zerolog.Level
	LoggerFile     bool
	GlobalCtx      string
)

func initOpenapiartlog() error {
	Openapiartlog = &openapiartlog{
		RootDir:       new(string),
		LogDir:        new(string),
		MaxLogSizeMB:  new(int),
		MaxLogBackups: new(int),
	}
	*Openapiartlog.MaxLogSizeMB = 25
	*Openapiartlog.MaxLogBackups = 5
	if *Openapiartlog.RootDir = os.Getenv("SRC_ROOT"); *Openapiartlog.RootDir == "" {
		*Openapiartlog.RootDir = "."
	}
	*Openapiartlog.LogDir = path.Join(*Openapiartlog.RootDir, "logs")
	GlobalLogger = nil
	GlobalLogLevel = zerolog.InfoLevel
	GlobalCtx = ""
	LoggerFile = false
	return nil
}

func SetUserLogger(logger zerolog.Logger) {
	GlobalLogger = &logger
}

func SetUserLogOutputToFile(choice bool) zerolog.Logger {
	LoggerFile = choice
	return GetLogger(GlobalCtx)
}

func SetUserLogLevel(logLevel zerolog.Level) {
	GlobalLogLevel = logLevel
	RefreshLogLevel(GlobalLogLevel)
}

func GetLogger(ctx string) zerolog.Logger {
	if GlobalCtx == "" {
		initOpenapiartlog()
	}
	GlobalCtx = ctx
	var localLogger zerolog.Logger
	if !LoggerFile {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		localLogger = zerolog.New(consoleWriter).Level(GlobalLogLevel).With().Timestamp().Str("Module", ctx).Logger()

	} else {
		if err := initFileLogger(); err != nil {
			panic(fmt.Errorf("Logger init failed: %v", err))
		}
		localLogger = log.With().Str("Module", ctx).Logger()
	}
	GlobalLogger = &localLogger
	return *GlobalLogger
}

func initLogger() (err error) {
	writer := lumberjack.Logger{
		Filename:   path.Join(*Openapiartlog.LogDir, "openapiartlog.log"),
		MaxSize:    *Openapiartlog.MaxLogSizeMB,
		MaxBackups: *Openapiartlog.MaxLogBackups,
	}
	zerolog.TimestampFieldName = "ts"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = zerolog.New(&writer).Level(GlobalLogLevel).With().Timestamp().Logger()
	return err
}

func RefreshLogLevel(logLevel zerolog.Level) {
	if !LoggerFile {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).Level(logLevel)
	} else {
		log.Logger = log.Level(logLevel)
	}
}

func initFileLogger() (err error) {
	if err = os.RemoveAll(*Openapiartlog.LogDir); err != nil {

		return err
	}
	if err = os.MkdirAll(*Openapiartlog.LogDir, os.ModePerm); err != nil {
		return err
	}
	if err := os.Chmod(*Openapiartlog.LogDir, 0771); err != nil {
		return err
	}
	return initLogger()
}
