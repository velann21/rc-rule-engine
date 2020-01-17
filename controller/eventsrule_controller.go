package controller

import (
	"context"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/entities/responses"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/models"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/service"
	"net/http"
)

func SyncAppsController(rw http.ResponseWriter, req *http.Request){
	eventsRequest := models.SyncApps{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateSyncApps(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateSyncAppsRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.SyncAppsEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}

func DeployAppsController(rw http.ResponseWriter, req *http.Request){
	eventsRequest := models.DeployApps{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeployApps(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeployAppsRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeployAppsEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}

func AddNodeController(rw http.ResponseWriter, req *http.Request) {
	eventsRequest := models.AddNode{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateAddNode(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateAddNodeRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.AddNodeEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}


func DeleteNodeController(rw http.ResponseWriter, req *http.Request) {
	eventsRequest := models.DeleteNode{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeleteNode(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeleteNodeRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeleteNodeEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}

func CreateClusterController(rw http.ResponseWriter, req *http.Request) {
	eventsRequest := models.CreateCluster{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateCreateCluster(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateCreateClusterRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.CreateClusterEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}

func DeleteClusterController(rw http.ResponseWriter, req *http.Request){
	eventsRequest := models.DeleteCluster{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeleteCluster(req.Body)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeleteClusterRequest()
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeleteClusterEvents(ctx, &eventsRequest)
	if err != nil {
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	return
}

func ReloadRuleSetController(rw http.ResponseWriter, req *http.Request){
	successResponse := responses.Response{}
    reloadRuleSet := models.ReloadRuleSet{}
	err := reloadRuleSet.PopulateRuleSet(req.Body)
	if err != nil{
		responses.HandleError(rw, err)
		return
	}
	reloadRuleSet.ValidateAndResetPath()
	ctx := context.Background()
    service.ReloadRuleSet(ctx, &reloadRuleSet)
	successResponse.SendResponse(rw, 200)
	return
}





