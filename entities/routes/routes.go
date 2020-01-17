package routes

import (
	"encoding/json"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
	)

type Events struct {
	EventType string `json:"eventType"`
}

//type DistributedTracingTransaction struct {
//	EventType string    `json:"eventType"`
//}

func (events *Events) PopulateDTEventsStruct(body []byte) (*Events, error) {
	err := json.Unmarshal(body, events)
	if err != nil {
		return nil, helpers.ErrInvalidRequest
	}
	return events, nil
}

