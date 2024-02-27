package monitor

import (
	"log"
	"net/http"
	"time"

	"github.com/stenstromen/linkvigil/api"
	"github.com/stenstromen/linkvigil/types"
)

func Monitor(endpoints []types.Endpoint, debug bool) {
	ticker := time.NewTicker(300 * time.Second)
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
					api.UpdateComponentStatus(http.DefaultClient, endpoint, types.MajorOutage)
				}
				continue
			}

			responseTime := time.Since(start)

			switch {
			case responseTime > 500*time.Millisecond:
				if debug {
					log.Printf("DEBUG WARNING - %s Degraded Performance", endpoint.Name)
				} else {
					log.Printf("WARNING - %s Degraded Performance", endpoint.Name)
					api.UpdateComponentStatus(http.DefaultClient, endpoint, types.DegradedPerformance)
				}
			case resp.StatusCode == 200:
				if debug {
					log.Printf("DEBUG OK - %s HTTP Status 200", endpoint.Name)
				} else {
					api.UpdateComponentStatus(http.DefaultClient, endpoint, types.Operational)
				}
			case resp.StatusCode == 500:
				if debug {
					log.Printf("DEBUG CRITICAL - %s Malfunctioning", endpoint.Name)
				} else {
					log.Printf("CRITICAL - %s Malfunctioning", endpoint.Name)
					api.UpdateComponentStatus(http.DefaultClient, endpoint, types.MajorOutage)
				}
			default:
				if debug {
					log.Printf("DEBUG INFO - %s HTTP Status %s", endpoint.Name, resp.Status)
				} else {
					log.Printf("INFO - %s HTTP Status %s", endpoint.Name, resp.Status)
				}
			}

		}
	}
}
