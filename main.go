package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var version = "0.0.1"

var (
	addr                  = flag.String("listen-address", ":5000", "The address to listen on.")
	configFile            = flag.String("config.file", "config.yml", "Path to configuration file.")
	debug                 = flag.Bool("debug", false, "Add verbose logging")
	showVersion           = flag.Bool("v", false, "prints current yace version.")
	cloudwatchConcurrency = flag.Int("cloudwatch-concurrency", 5, "Maximum number of concurrent requests to CloudWatch API")
	tagConcurrency        = flag.Int("tag-concurrency", 5, "Maximum number of concurrent requests to Resource Tagging API")

	supportedServices = []string{
		"ec",
	}

	config = conf{}
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	log.Println("Parse config..")
	if err := config.load(configFile); err != nil {
		log.Fatal("Couldn't read ", *configFile, ":", err)
	}

	cloudwatchSemaphore = make(chan struct{}, *cloudwatchConcurrency)
	tagSemaphore = make(chan struct{}, *tagConcurrency)
	RunForever()
}
