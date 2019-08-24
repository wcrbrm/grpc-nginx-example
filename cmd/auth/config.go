package main

import (
	cli "github.com/jawher/mow.cli"
)

var (
	appName = app.String(cli.StringOpt{
		Name:   "name",
		Desc:   "",
		EnvVar: "APP_NAME",
		Value:  "auth",
	})
	envName = app.String(cli.StringOpt{
		Name:   "env",
		Desc:   "",
		EnvVar: "APP_ENV",
		Value:  "local",
	})
	appLogLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "APP_LOG_LEVEL",
		Value:  "info",
	})
	appDbConnection = app.String(cli.StringOpt{
		Name:   "dsn",
		Desc:   "Specify database connection string.",
		EnvVar: "PGSQL_DSN",
		Value:  "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable",
	})
	appGrpcPort = app.Int(cli.IntOpt{
		Name:   "p port",
		Desc:   "",
		EnvVar: "APP_PORT",
		Value:  50051,
	})
	
)
