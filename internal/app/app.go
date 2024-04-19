package app

import (
	"context"
	"fmt"
	"halalapi/config"
	"halalapi/pkg/datasource"
	"halalapi/pkg/otel/zerolog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/gommon/log"
)

type App struct {
	db  *mongo.Client
	gin *gin.Engine
	log *zerolog.Logger
	cfg config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	return &App{
		db:  db,
		gin: gin.Default(),
		log: zerolog.NewZeroLog(ctx, os.Stdout),
		cfg: cfg,
	}
}

func (app *App) Run() error {
	if err := app.startService(); err != nil {
		app.log.Z().Err(err).Msg("[app]StartService")

		return err
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit
		log.Info("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		app.db.Disconnect(ctx)
	}()

	// for create middleware

	return app.gin.Run(fmt.Sprintf(":%s", app.cfg.Server.RESTPort))
}
