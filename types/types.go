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
		Description        string          `json:"description"`
		Status             ComponentStatus `json:"status"`
		Name               string          `json:"name"`
		OnlyShowIfDegraded bool            `json:"only_show_if_degraded"`
		GroupId            string          `json:"group_id"`
		Showcase           bool            `json:"showcase"`
		StartDate          string          `json:"start_date"`
	} `json:"component"`
}
