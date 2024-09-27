package app

import (
	"context"
	"github.com/wasd0/is/pkg/app"
	"github.com/wasd0/is/pkg/config"
	"github.com/wasd0/is/pkg/logger"
	"github.com/wasd0/is/pkg/logger/zero"
	"is-tgbot/internal/bot"
	"os/signal"
	"syscall"
)

func Startup() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	defer stop()
	runServer(ctx)
}

func runServer(ctx context.Context) {
	cfg := config.MustRead()
	closer := &app.Closer{}
	_, loggerCallback := zero.MustSetUp(cfg)

	closer.Add(loggerCallback)

	printStartMessage(cfg)

	bot.Start(ctx)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		logger.Log().Fatal(err, "Server close failed")
	}
}

func printStartMessage(cfg *config.Config) {
	logger.Log().Info("Telegram bot  started")
	logger.Log().Infof("Host: %s", cfg.Server.Host)
	logger.Log().Infof("Port: %s", cfg.Server.Port)
	logger.Log().Infof("ENV: %s", cfg.Env)
}