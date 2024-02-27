package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/stenstromen/linkvigil/types"
)

func UpdateComponentStatus(c *http.Client, endpoint types.Endpoint, status types.ComponentStatus) {
	updateRequest := types.ComponentUpdateRequest{}
	updateRequest.Component.Status = status

	reqBody, err := json.Marshal(updateRequest)
	if err != nil {
		log.Fatalf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("PATCH", "https://api.statuspage.io/v1/pages/"+endpoint.PageID+"/components/"+endpoint.ComponentID, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "OAuth "+os.Getenv("STATUSPAGE_API_KEY"))
	req.Header.Set("User-Agent", "LinkVigil/1.0")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error executing request: %v", err)
	}

	defer resp.Body.Close()
}
