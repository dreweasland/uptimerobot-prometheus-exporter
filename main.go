package main

import (
	"log"
	"os"

	"github.com/dreweasland/uptimerobot-prometheus-exporter/cmd"
)

var (
	logLevel string
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
