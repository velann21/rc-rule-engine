package requests

import (
	"encoding/json"
	"io"
	"log"

	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
)

// EventsRequest is ...
type EventsRequest struct {
	EventType    string `json:"eventType"`
	EventOccured int    `json:"eventOccured"`
}

func (eventsRequest *EventsRequest) PopulateEventsRequest(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *EventsRequest) ValidateEventsRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	return helpers.ErrInvalidRequest
}
