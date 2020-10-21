package main

type AccountAggregationSchedule struct {
	Arguments struct {
		Method  string `json:"method,omitempty"`
		Path    string `json:"path,omitempty"`
		Service string `json:"service,omitempty"`
	} `json:"arguments"`
	SourceID        string      `json:"source_id,omitempty"`
	CronExpressions []string    `json:"cronExpressions,omitempty"`
	Description     string      `json:"description,omitempty"`
	LastExecution   interface{} `json:"lastExecution,omitempty"`
	Name            string      `json:"name,omitempty"`
	NewState        interface{} `json:"newState,omitempty"`
	NextExecution   int64       `json:"nextExecution,omitempty"`
	State           interface{} `json:"state,omitempty"`
	Type            string      `json:"type,omitempty"`
}
