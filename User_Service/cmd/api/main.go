package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	restapi "github.com/DiMashina05/FooDiMashina/User_Service/internal/rest_api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := restapi.NewHandler(logger)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	serveChan := make(chan error, 1)
	go func() {
		serveChan <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		logger.Info("Завершение работы командой")
		if err := srv.Shutdown(ctx); err != nil {

			logger.Error("Ошибка во время выключения",
				"err", err,
				"addr", srv.Addr,
			)
		}
	case err := <-serveChan:

		logger.Error("ошибка сервера",
			"err", err,
			"addr", srv.Addr,
		)
	}
}
