package types

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
	MajorOutage         ComponentStatus = "major_outage"
)

type ComponentUpdateRequest struct {
	Component struct {
		Status ComponentStatus `json:"status"`
	} `json:"component"`
}
