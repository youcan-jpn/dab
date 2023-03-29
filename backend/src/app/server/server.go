package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
	"github.com/youcan-jpn/dab/backend/src/app/api"
)

type HTTPServerOption struct {
	Address     string
	Port        int
	DBClient    *sql.DB
	Environment string
}

func RunHTTPServer(ctx context.Context, opt HTTPServerOption) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	e := echo.New()

	timezone, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	apiServer := api.NewAPI(
		ctx,
		opt.DBClient,
		opt.Environment,
		timezone,
	)

	oapi.RegisterHandlersWithBaseURL(e, apiServer, "/api")

	s := http.Server{
		Addr:    fmt.Sprintf("%s:%d", opt.Address, opt.Port),
		Handler: e,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
