package main

import (
	"context"
	"github.com/data-preservation-programs/singularity/cmd"
	log2 "github.com/ipfs/go-log/v2"
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
	if log2.GetConfig().Level > log2.LevelInfo && os.Getenv("GOLOG_LOG_LEVEL") == "info" {
		log2.SetAllLoggers(log2.LevelInfo)
	}
	if err := cmd.RunApp(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
