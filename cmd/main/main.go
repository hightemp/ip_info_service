package main

import (
	"flag"

	"github.com/hightemp/ip_info_service/internal/config"
	"github.com/hightemp/ip_info_service/internal/data_parser"
	"github.com/hightemp/ip_info_service/internal/logger"
	"github.com/hightemp/ip_info_service/internal/models/ip_range"
	"github.com/hightemp/ip_info_service/internal/server"
)

func init() {
	logger.InitLogger()
}

func main() {
	cfg := flag.String("c", "config.yaml", "configuration file")
	flag.Parse()

	err := config.Load(*cfg)

	if err != nil {
		logger.Panic("%v", err)
	}

	logger.LogInfo("Trying to load data from json")
	err = data_parser.Load()

	if err != nil {
		logger.Panic("Can't load data: %v", err)
	}

	logger.LogInfo("Trying to load data from yaml")
	err = ip_range.Load()

	if err != nil {
		logger.Panic("Can't load data: %v", err)
	}

	server.Start()
}
