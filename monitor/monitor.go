package monitor

import (
	"log"
	"net/http"
	"time"

	"github.com/stenstromen/linkvigil/types"
)

func Monitor(endpoints []types.Endpoint, debug bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, endpoint := range endpoints {
			start := time.Now()
			resp, err := http.Get(endpoint.URL)
			if err != nil {
				if debug {
					log.Printf("DEBUG CRITICAL - %s HTTP request failed: %v", endpoint.Name, err)
				} else {
					log.Printf("CRITICAL - %s HTTP request failed: %v", endpoint.Name, err)
				}
				continue
			}

			responseTime := time.Since(start)

			if responseTime > 500*time.Millisecond {
				if debug {
					log.Printf("DEBUG WARNING - %s Degraded Performance", endpoint.Name)
				} else {
					log.Printf("WARNING - %s Degraded Performance", endpoint.Name)
				}
			}

			if resp.StatusCode == 200 {
				if debug {
					log.Printf("DEBUG OK - %s HTTP Status 200", endpoint.Name)
				} else {
					log.Printf("OK - %s HTTP Status 200", endpoint.Name)
				}
			} else if resp.StatusCode == 500 {
				if debug {
					log.Printf("DEBUG CRITICAL - %s Malfunctioning", endpoint.Name)
				} else {
					log.Printf("CRITICAL - %s Malfunctioning", endpoint.Name)
				}
			} else {
				if debug {
					log.Printf("DEBUG INFO - %s HTTP Status %s", endpoint.Name, resp.Status)
				} else {
					log.Printf("INFO - %s HTTP Status %s", endpoint.Name, resp.Status)
				}
			}
		}
	}
}
