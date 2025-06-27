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

//go:generate sh ./docgen.sh
//go:generate go run github.com/swaggo/swag/cmd/swag@v1.16.3 init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./docs/swagger
//go:generate rm -rf ./docs/en/web-api-reference
//go:generate go run docs/gen/webapireference/main.go
//go:generate rm -rf ./client
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0 generate client -f ./docs/swagger/swagger.json -t . -c client/swagger/http -m client/swagger/models -a client/swagger/operations -q

//go:embed version.json
var versionJSON []byte

func init() {
	if os.Getenv("GOLOG_LOG_LEVEL") == "" {
		_ = os.Setenv("GOLOG_LOG_LEVEL", "info")
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
// @contact.name   Singularity Team
// @contact.url    https://github.com/data-preservation-programs/singularity/issues
// @license.name MIT + Apache 2.0
// @license.url https://github.com/data-preservation-programs/singularity/blob/main/LICENSE
// @accept json
// @produce json
func main() {
	if log2.GetConfig().Level > log2.LevelInfo && os.Getenv("GOLOG_LOG_LEVEL") == "info" {
		log2.SetAllLoggers(log2.LevelInfo)
	}
	cmd.SetupHelpPager()
	cmd.SetupErrorHandler()
	err := cmd.SetVersionJSON(versionJSON)
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.App.RunContext(context.TODO(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
