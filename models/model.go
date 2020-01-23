package models

import (
	"encoding/json"
	eventsLogger "gitlab.reynencourt.com/reynen-court/rc-common-lib/events_logger"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
	"io"
	"log"
	"os"
)

const (
	SyncAppFailed = "SyncAppFailed"
	DeployAppFailed = "DeployAppFailed"
	AddNodeFailed = "AddNodeFailed"
	DeleteNodeFailed = "DeleteNodeFailed"
	CreateClusterFailed = "CreateClusterFailed"
	DeleteClusterFailed = "DeleteClusterFailed"
)

type SyncApps struct {
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
	Units string `json:"units"`
}


func (eventsRequest *SyncApps) PopulateSyncApps(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	eventsRequest.Operator = "na"
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *SyncApps) MarshalSyncAppsEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *SyncApps) ValidateSyncAppsRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ErrorCode != "" {
	   if eventsRequest.ServiceName == eventsLogger.SvcDeploymentManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
		eventsRequest.EventType = SyncAppFailed
	}else{
		return helpers.ErrInvalidRequest
	   }
	}
	return nil
}

func (p *SyncApps) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *SyncApps) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.TimeFrame = timeFrame
	p.Units = units
	return true
}


type DeployApps struct {
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
	Units string `json:"units"`
}


func (eventsRequest *DeployApps) PopulateDeployApps(body io.ReadCloser) error {
	eventsRequest.Operator = "na"
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *DeployApps) MarshalDeployAppsEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *DeployApps) ValidateDeployAppsRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ErrorCode != "" {
		if eventsRequest.ServiceName == eventsLogger.SvcDeploymentManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
			eventsRequest.EventType = DeployAppFailed
		}else{
			return helpers.ErrInvalidRequest
		}
	}

	return nil

}

func (p *DeployApps) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *DeployApps) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.TimeFrame = timeFrame
	p.Units = units
	return true
}


type CreateCluster struct {
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
	Units string `json:"units"`
}


func (eventsRequest *CreateCluster) PopulateCreateCluster(body io.ReadCloser) error {
	eventsRequest.Operator = "na"
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *CreateCluster) MarshalCreateClusterEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *CreateCluster) ValidateCreateClusterRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ErrorCode != "" {
		if eventsRequest.ServiceName == eventsLogger.SvcResourceManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
			eventsRequest.EventType = CreateClusterFailed
		}else{
			return helpers.ErrInvalidRequest
		}
	}
	return nil
}

func (p *CreateCluster) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *CreateCluster) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.TimeFrame = timeFrame
	p.Units = units
	return true
}


type DeleteCluster struct {
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
	Units string `json:"units"`
}


func (eventsRequest *DeleteCluster) PopulateDeleteCluster(body io.ReadCloser) error {
	eventsRequest.Operator = "na"
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *DeleteCluster) MarshalDeleteClusterEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *DeleteCluster) ValidateDeleteClusterRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" {
		if eventsRequest.ServiceName == eventsLogger.SvcResourceManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
			eventsRequest.EventType = DeleteClusterFailed
		}else{
			return helpers.ErrInvalidRequest
		}
	}
    return nil
}

func (p *DeleteCluster) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *DeleteCluster) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.TimeFrame = timeFrame
	p.Units = units
	return true
}

type AddNode struct {
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
	Units string `json:"units"`
}


func (eventsRequest *AddNode) PopulateAddNode(body io.ReadCloser) error {
	eventsRequest.Operator = "na"
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *AddNode) MarshalAddNodeEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *AddNode) ValidateAddNodeRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" {
		if eventsRequest.ServiceName == eventsLogger.SvcResourceManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
			eventsRequest.EventType = AddNodeFailed
		}else{
			return helpers.ErrInvalidRequest
		}
	}
	return nil

}

func (p *AddNode) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *AddNode) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.Units = units
	p.TimeFrame = timeFrame
	return true
}


type DeleteNode struct {
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
	Units string `json:"units"`
}


func (eventsRequest *DeleteNode) PopulateDeleteNode(body io.ReadCloser) error {
	eventsRequest.Operator = "na"
	decoder := json.NewDecoder(body)
	err := decoder.Decode(eventsRequest)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (eventsRequest *DeleteNode) MarshalDeleteNodeEvents()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(eventsRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

func (eventsRequest *DeleteNode) ValidateDeleteNodeRequest() error {
	if eventsRequest.EventType == "" {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.EventOccured == 0 {
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ServiceName == "" || eventsRequest.ServiceName == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.DateTime == "" || eventsRequest.DateTime == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ActionType == "" || eventsRequest.ActionType == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.TraceID == "" || eventsRequest.TraceID == " "{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" && eventsRequest.ActionType == "Request"{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode == "" {
		if eventsRequest.ActionType == "Response"{
			return helpers.ErrInvalidRequest
		}

	}
	if eventsRequest.ActionType == "Request" && eventsRequest.ServiceName != eventsLogger.SvcControllerBackend{
		return helpers.ErrInvalidRequest
	}
	if eventsRequest.ErrorCode != "" {
		if eventsRequest.ServiceName == eventsLogger.SvcResourceManager || eventsRequest.ServiceName == eventsLogger.SvcControllerBackend{
			eventsRequest.EventType = DeleteNodeFailed
		}else{
			return helpers.ErrInvalidRequest
		}
	}
	return nil

}

func (p *DeleteNode) SetAlertType(alertType string, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.AlertType = alertType
	return true
}

func (p *DeleteNode) SetTimeFrame(operator string,units string, timeFrame float64, prevVal ...bool) bool {
	if len(prevVal) > 0 {
		if !prevVal[0] {
			return false
		}
	}
	p.Operator = operator
	p.TimeFrame = timeFrame
	p.Units = units
	return true
}

type ReloadRuleSet struct {
	FilePath string
}

func (reloadRuleSet *ReloadRuleSet) PopulateRuleSet(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(reloadRuleSet)
	if err != nil {
		log.Println("Error While Populate the SyncApps struct")
		return helpers.ErrInvalidRequest
	}
	return nil
}

func (reloadRuleSet *ReloadRuleSet) ValidateAndResetPath(){
	if reloadRuleSet.FilePath == ""{
		//_, b, _, _ := runtime.Caller(0)
		//basepath   := filepath.Dir(b)
       reloadRuleSet.FilePath = os.Getenv("RULESET_FILEPATH")
	}
}





