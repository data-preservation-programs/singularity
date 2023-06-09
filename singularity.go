package main

import (
	"context"
	"github.com/data-preservation-programs/singularity/cmd"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

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
func main() {
	if err := cmd.RunApp(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
