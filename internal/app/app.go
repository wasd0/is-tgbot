package app

import (
	"context"
	"github.com/wasd0/is-common/pkg/app"
	"github.com/wasd0/is-common/pkg/config"
	"github.com/wasd0/is-common/pkg/logger"
	"github.com/wasd0/is-common/pkg/logger/zero"
	"is-tgbot/internal/app/serviceProvider"
	"is-tgbot/internal/bot"
	"is-tgbot/internal/command"
	"is-tgbot/internal/storage"
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
	_, loggerCallback := zero.MustSetUp(cfg)
	redisClient, redisCallback := storage.MustOpenRedis(ctx)
	closer := &app.Closer{}

	servProvider := serviceProvider.NewServiceProvider(redisClient)
	commandProvider := command.NewCommandProvider(servProvider)

	closer.Add(loggerCallback)
	closer.Add(redisCallback)

	printStartMessage(cfg)

	bot.Start(ctx, commandProvider)

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
