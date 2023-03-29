package api

import (
	"context"
	"database/sql"
	"time"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
)

type API struct {
	oapi.ServerInterface // 埋め込む必要なさそう
	ctx                  context.Context
	dbClient             *sql.DB
	envString            string
	timezone             *time.Location
}

func NewAPI(
	ctx context.Context,
	dbClient *sql.DB,
	env string,
	timezone *time.Location,
) *API {
	return &API{
		ctx:       ctx,
		dbClient:  dbClient,
		envString: env,
		timezone:  timezone,
	}
}
