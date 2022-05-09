package main

import (
	"github.com/anik3tra0/guidecx-cli/cmd"
	"log"
	"os"
	"runtime"
)

var (
	// VERSION Version of CLI
	VERSION      = "0.0.1"
	VersionLabel = "alpha"
)

func main() {
	log.Printf(
		"You are running %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)
	_, present := os.LookupEnv("GUIDECX_API_KEY")
	if !present {
		log.Println("GUIDECX_API_KEY Environment Variable Not Set")
		os.Exit(1)
	}

	cmd.Execute()
}
