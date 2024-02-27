package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Endpoint struct {
	Name        string `yaml:"name"`
	URL         string `yaml:"url"`
	PageID      string `yaml:"page_id"`
	ComponentID string `yaml:"component_id"`
}

type ComponentStatus string

const (
	Operational         ComponentStatus = "operational"
	DegradedPerformance ComponentStatus = "degraded_performance"
	PartialOutage       ComponentStatus = "partial_outage"
	MajorOutage         ComponentStatus = "major_outage"
)

type ComponentUpdateRequest struct {
	Component struct {
		Description        string          `json:"description"`
		Status             ComponentStatus `json:"status"`
		Name               string          `json:"name"`
		OnlyShowIfDegraded bool            `json:"only_show_if_degraded"`
		GroupId            string          `json:"group_id"`
		Showcase           bool            `json:"showcase"`
		StartDate          string          `json:"start_date"`
	} `json:"component"`
}

func updateComponentStatus(c *http.Client, endpoint Endpoint, status ComponentStatus) {
	updateRequest := ComponentUpdateRequest{}
	updateRequest.Component.Status = status
	updateRequest.Component.Description = "Your description here"
	updateRequest.Component.Name = "Your component name here"
	updateRequest.Component.OnlyShowIfDegraded = true
	updateRequest.Component.GroupId = "Your group ID here"
	updateRequest.Component.Showcase = true
	updateRequest.Component.StartDate = "2024-02-24"

	reqBody, err := json.Marshal(updateRequest)
	if err != nil {
		log.Fatalf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("PATCH", "https://api.statuspage.io/v1/pages/"+endpoint.PageID+"/components/"+endpoint.ComponentID, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "OAuth "+os.Getenv("STATUSPAGE_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error executing request: %v", err)
	}
	defer resp.Body.Close()
}

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

	var endpoints []Endpoint
	err = yaml.Unmarshal([]byte(data), &endpoints)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

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
