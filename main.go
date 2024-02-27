package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/stenstromen/linkvigil/monitor"
	"github.com/stenstromen/linkvigil/types"
)

func main() {
	var endpointsFilePath string
	var debug bool = false

	args := os.Args[1:]
	if len(args) > 0 {
		endpointsFilePath = args[0]
		if len(args) > 1 && args[1] == "debug" {
			debug = true
		}
	} else {
		log.Fatalf("error: no endpoints file provided")
	}

	data, err := os.ReadFile(endpointsFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var endpoints []types.Endpoint
	err = yaml.Unmarshal([]byte(data), &endpoints)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	monitor.Monitor(endpoints, debug)
}
