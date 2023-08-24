package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	"github.com/data-preservation-programs/singularity/cmd"
	log2 "github.com/ipfs/go-log/v2"
	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run handler/storage/create/gen/main.go
//go:generate sh ./docgen.sh
//go:generate go run github.com/swaggo/swag/cmd/swag@v1.8.12 init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./docs/swagger
//go:generate go run docs/gen/webapireference/main.go

//go:embed version.json
var version []byte

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
	cmd.SetupHelpPager()
	err := cmd.SetVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.App.RunContext(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
