package main

import (
	"context"
	"halalapi/config"
	"halalapi/internal/app"
	"os"
)

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	app.NewApp(context.Background(), cfg).Run()
}
