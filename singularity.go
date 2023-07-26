package main

import (
	"context"
	"log"
	"os"

	"github.com/data-preservation-programs/singularity/cmd"
	log2 "github.com/ipfs/go-log/v2"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run github.com/swaggo/swag/cmd/swag@v1.8.12 init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./api/docs

func init() {
	if os.Getenv("GOLOG_LOG_LEVEL") == "" {
		os.Setenv("GOLOG_LOG_LEVEL", "info")
	}
}

// @title Singularity API
// @version beta
// @description This is the API for Singularity, a tool for large-scale clients with PB-scale data onboarding to Filecoin network.
// @host localhost:9090
// @BasePath /api
// @securityDefinitions none
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @contact.name   Xinan Xu
// @contact.url    https://github.com/data-preservation-programs/singularity/issues
// @license.name MIT + Apache 2.0
// @license.url https://github.com/data-preservation-programs/singularity/blob/main/LICENSE
func main() {
	if log2.GetConfig().Level > log2.LevelInfo && os.Getenv("GOLOG_LOG_LEVEL") == "info" {
		log2.SetAllLoggers(log2.LevelInfo)
	}
	if err := cmd.RunApp(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
