package service

import (
	"context"
	"fmt"
	"github.com/myntra/roulette/log"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/dao"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/database"
	dbModel "gitlab.reynencourt.com/reynen-court/rc-rules-engine/entities/database_models"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
	requests "gitlab.reynencourt.com/reynen-court/rc-rules-engine/models"
	"time"
)

func SyncAppsEvents(ctx context.Context, eventsRequest *requests.SyncApps) error {
	ruleSetObj := helpers.GetRuleSetObject()
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}

func DeployAppsEvents(ctx context.Context, eventsRequest *requests.DeployApps) error {
	ruleSetObj := helpers.GetRuleSetObject()
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}


func AddNodeEvents(ctx context.Context, eventsRequest *requests.AddNode) error {
	ruleSetObj := helpers.GetRuleSetObject()
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}

func DeleteNodeEvents(ctx context.Context, eventsRequest *requests.DeleteNode) error {
	ruleSetObj := helpers.GetRuleSetObject()
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}


func CreateClusterEvents(ctx context.Context, eventsRequest *requests.CreateCluster) error {
	ruleSetObj := helpers.GetRuleSetObject()
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}

func DeleteClusterEvents(ctx context.Context, eventsRequest *requests.DeleteCluster) error {
	ruleSetObj := helpers.GetRuleSetObject()
	fmt.Println(ruleSetObj)
	etcdConnection := database.GetEtcdConnection()
	resp, err := dao.GetEvent(ctx, etcdConnection, eventsRequest.EventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	events := dbModel.Events{}
	if resp.Kvs != nil{
		events.PopulateEvents(resp.Kvs[0].Value)
		eventsRequest.EventOccured = eventsRequest.EventOccured+events.EventOccured
		eventsRequest.LastUpdated = events.LastUpdated
	}

	if resp.Kvs == nil{
		eventsRequest.LastUpdated = eventsRequest.DateTime
	}

	ruleSetObj.Executor.Execute(eventsRequest)
	go takeEventForNotification(eventsRequest)
	return nil
}

func ReloadRuleSet(ctx context.Context, eventsRequest *requests.ReloadRuleSet){
	ruleSetObj := helpers.GetRuleSetObject()
	ruleSetObj.LoadRuleSet(eventsRequest.FilePath)
}

var syncAppsChan =  make(chan *requests.SyncApps, 100)
var deployAppsChan =  make(chan *requests.DeployApps, 100)
var addNodeChan =  make(chan *requests.AddNode, 100)
var deleteNodeChan =  make(chan *requests.DeleteNode, 100)
var createClusterChan =  make(chan *requests.CreateCluster, 100)
var deleteClusterChan =  make(chan *requests.DeleteCluster, 100)
func takeEventForNotification(events interface{}){
	switch events.(type) {
	case *requests.SyncApps:
		syncAppsChan <- events.(*requests.SyncApps)
	case *requests.DeployApps:
		deployAppsChan <- events.(*requests.DeployApps)
	case *requests.AddNode:
		addNodeChan <- events.(*requests.AddNode)
	case *requests.DeleteNode:
		deleteNodeChan <- events.(*requests.DeleteNode)
	case *requests.CreateCluster:
		createClusterChan <- events.(*requests.CreateCluster)
	case *requests.DeleteCluster:
		deleteClusterChan <- events.(*requests.DeleteCluster)
	}
}


func ExecuteEventForNotification(){
	for {
		select {
		case task := <-syncAppsChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalSyncAppsEvents()
				if err != nil{
					return
				}
				if shouldAlertTriggered{
					alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}
					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{

						"alertname":"testalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName, "severity":"warning"})
				}
				if isDeleteApproved{
					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}

				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}

				return
			}()
		case task := <-deployAppsChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalDeployAppsEvents()
				if err != nil{
					return
				}
				if shouldAlertTriggered{
					alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}
					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
						"alertname":"testalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName,"severity":"warning"})
				}
				if isDeleteApproved{
					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}
				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}
				return
			}()
		case task := <-addNodeChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalAddNodeEvents()
				if err != nil{
					return
				}
				fmt.Println("IsDelete Approved", isDeleteApproved)
				fmt.Println(shouldAlertTriggered)
				if shouldAlertTriggered{
					fmt.Println("Sending an Email")
					alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}
					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
						"alertname":"testalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName, "severity":"warning"})
				}
				if isDeleteApproved{
					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}
				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}
				return
			}()
		case task := <-deleteNodeChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalDeleteNodeEvents()
				if err != nil{
					return
				}
				if shouldAlertTriggered{
					alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}
					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
						"alertname":"testalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName, "severity":"warning"})
				}
				fmt.Println("IsDelete Approved", isDeleteApproved)
				fmt.Println(shouldAlertTriggered)
				if isDeleteApproved{

					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}
				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}
				return
			}()
		case task := <-createClusterChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalCreateClusterEvents()
				if err != nil{
					return
				}
				if shouldAlertTriggered{
					alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}
					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
						"alertname":"testalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName, "severity":"warning"})
				}
				if isDeleteApproved{
					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}
				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}
				return
			}()
		case task := <-deleteClusterChan:
			go func(){
				etcdConnection := database.GetEtcdConnection()
				isDeleteApproved, shouldAlertTriggered := checkTimeFrameAndSendNotification(task.Units,task.TimeFrame,task.Operator,
					task.LastUpdated,task.DateTime,task.AlertType)
				task.LastUpdated = task.DateTime
				jsonData, err := task.MarshalDeleteClusterEvents()
				if err != nil{
					return
				}
				if shouldAlertTriggered{
                    alertManagerFactoryProd := helpers.AlertManagerFactoryProducer{}

					alertManagerFactoryProd.GetAlertManagerFactory(helpers.PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
						"alertname":"eventsbasedalert","EventsOccured":string(task.EventOccured),"ErrorCode":task.ErrorCode,"TraceID":task.TraceID,"EventType":task.EventType,"ServiceName":task.ServiceName,
						"severity":"warning"})
				}
				if isDeleteApproved{
					err = dao.DeleteEvent(context.Background(), etcdConnection, task.EventType)
					if err != nil{
						return
					}
					return
				}
				err = dao.StoreEvent(context.Background(), etcdConnection,task.EventType, jsonData)
				if err != nil{
					return
				}
				return
			}()
		}
	}
}

func checkTimeFrameAndSendNotification(unit string,timeFrame float64, operator string,
	lastTime string, currentTime string, alertType string) (bool,bool) {
	if alertType == "" || alertType == " "{
		return false, false
	}
    if unit == "Minutes"{
		if timeFrame > 0{
			if operator == "gt"{
				fmt.Println("Minutes gt")
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() > timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "lt"{
				fmt.Println("Minutes lt")
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() < timeFrame{
					log.Info("Condition Statified to trigger an Email")
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "gte"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() >= timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "lte"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() <= timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "ne"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() != timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "eq"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Minutes() == timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "na"{
				return true, true
			}
		}

	} else if unit == "Hours"{
		if timeFrame > 0{
			if operator == "gt"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() > timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "lt"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() < timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "gte"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() >= timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "lte"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() <= timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "ne"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() != timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "eq"{
				lastUpdatedTime, err := time.Parse(time.RFC3339, lastTime)
				currentTime, err := time.Parse(time.RFC3339, currentTime)
				if err != nil{
					return false, false
				}
				delta := currentTime.Sub(lastUpdatedTime)
				if delta.Hours() == timeFrame{
					return false, true
				}else{
					return true, false
				}
			}
			if operator == "na"{
				fmt.Println("Am here at na")
				return true, true
			}
		}
	}
    return false, false
}


