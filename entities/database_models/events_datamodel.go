package database_models

import "encoding/json"

type Events struct {
	EventType    string `json:"eventType"`
	EventOccured int    `json:"eventOccured"`
	TimeFrame float64 `json:"timeFrame"`
	LastUpdated string `json:"lastUpdated"`
	AlertType string `json:"alertType"`
	ServiceName                string                     `json:"serviceName"`
	DateTime                   string                     `json:"time"`
	ActionType                 string                     `json:"actionType"`
	TraceID                    string                     `json:"traceID"`
	ErrorCode                  string                     `json:"errorCode"`
	Operator string `json:"operator"`
}

func (events *Events) PopulateEvents(data []byte)*Events{
	err := json.Unmarshal(data, events)
	if err != nil{
       return nil
	}
	return events

}
