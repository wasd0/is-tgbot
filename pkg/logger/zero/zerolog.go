package zero

import (
	"context"
	"fmt"
	"io"
	"is-tgbot/pkg/app"
	"is-tgbot/pkg/config"
	"is-tgbot/pkg/logger"
	"is-tgbot/pkg/utils"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Zerolog struct {
	log *zerolog.Logger
}

func (logger *Zerolog) Info(msg string) {
	logger.log.Info().Msg(msg)
}

func (logger *Zerolog) Infof(format string, args ...interface{}) {
	logger.log.Info().Msgf(format, args...)
}

func (logger *Zerolog) Fatal(err error, msg string) {
	if err == nil {
		return
	}
	logger.log.Fatal().Err(err).Msg(msg)
}

func (logger *Zerolog) Fatalf(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	if args == nil {
		args = []interface{}{}
	}

	logger.log.Fatal().Err(err).Msgf(format, args)
}

func (logger *Zerolog) Error(err error, msg string) {
	if err == nil {
		return
	}
	logger.log.Err(err).Msg(msg)
}

func (logger *Zerolog) Errorf(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	if args == nil {
		args = []interface{}{}
	}

	logger.log.Err(err).Msgf(format, args)
}

// MustSetUp setups zerolog. Returns zerolog object and callback function
func MustSetUp(cfg *config.Config) (*Zerolog, app.Callback) {
	var (
		level zerolog.Level
		file  *os.File
	)

	consoleWriter := getConsoleWriter()
	output := consoleWriter

	switch cfg.Env {
	case config.EnvDev, config.EnvStage:
		level = zerolog.TraceLevel
	case config.EnvProd:
		level = zerolog.ErrorLevel
		file = initOutputFile(cfg)
		output = zerolog.MultiLevelWriter(consoleWriter, file)
	}

	zerolog.SetGlobalLevel(level)

	zeroLogger := zerolog.New(output).With().Stack().CallerWithSkipFrameCount(3).Timestamp().Logger()
	zero := &Zerolog{log: &zeroLogger}

	logger.Setup(zero)

	return zero, func(ctx context.Context) error {
		zero.Info("Log file closing...")

		if file != nil {
			return utils.CloseFile(file)
		}

		return nil
	}
}

func getConsoleWriter() io.Writer {
	return zerolog.ConsoleWriter{
		Out:          os.Stdout,
		NoColor:      false,
		TimeFormat:   time.RFC850,
		TimeLocation: time.UTC,
	}
}

func initOutputFile(cfg *config.Config) *os.File {
	if len(cfg.LogPath) == 0 {
		panic("No log file path")
	}

	logFileName := time.Now().Format("2006-01-02.log")
	logFilePath, err := utils.CreateDirectoriesIfNotExists(cfg.LogPath)

	if err != nil {
		panic(fmt.Errorf("errorUtils while creating log file directories by path: %v", cfg.LogPath))
	}

	logFilePath.WriteString(utils.Separator)
	logFilePath.WriteString(logFileName)

	file, openFileErr := os.OpenFile(
		logFilePath.String(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if openFileErr != nil {
		panic(fmt.Errorf("errorUtils opening file: %v", file))
	}

	return file
}
