package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/api"
	environment "github.com/youcan-jpn/dab/backend/src/env"
	_ "github.com/youcan-jpn/dab/backend/src/logger"
)

const version = "0.0.0"

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run server. err:%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 環境変数の読み込み
	env, err := environment.Process()
	if err != nil {
		return err
	}

	// DBの設定
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName))
	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.CORS())

	timezone, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	apiServer := api.NewAPI(
		ctx,
		db,
		env.DabEnvironment,
		timezone,
	)

	oapi.RegisterHandlersWithBaseURL(e, apiServer, "/api")

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return err
}
