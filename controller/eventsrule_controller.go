package controller

import (
	"context"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/entities/responses"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/models"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/service"
	"log"
	"net/http"
)

func SyncAppsController(rw http.ResponseWriter, req *http.Request){
	log.Println("Inside The SyncAppsController in RE")
	eventsRequest := models.SyncApps{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateSyncApps(req.Body)
	if err != nil {
		log.Println("Error at The SyncAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateSyncAppsRequest()
	if err != nil {
		log.Println("Error at The SyncAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.SyncAppsEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The SyncAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Done The SyncAppsController in RE")
	return
}

func DeployAppsController(rw http.ResponseWriter, req *http.Request){
	log.Println("Inside The DeployAppsController in RE")
	eventsRequest := models.DeployApps{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeployApps(req.Body)
	if err != nil {
		log.Println("Error at The DeployAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeployAppsRequest()
	if err != nil {
		log.Println("Error at The DeployAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeployAppsEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The DeployAppsController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Done The DeployAppsController in RE")
	return
}

func AddNodeController(rw http.ResponseWriter, req *http.Request) {
	log.Println("Inside The AddNodeController in RE")
	eventsRequest := models.AddNode{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateAddNode(req.Body)
	if err != nil {
		log.Println("Error at The AddNodeController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateAddNodeRequest()
	if err != nil {
		log.Println("Error at The AddNodeController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.AddNodeEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The AddNodeController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Error at The AddNodeController in RE")
	return
}


func DeleteNodeController(rw http.ResponseWriter, req *http.Request) {
	log.Println("Inside The DeleteNodeController in RE")
	eventsRequest := models.DeleteNode{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeleteNode(req.Body)
	if err != nil {
		log.Println("Error at The DeleteNodeController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeleteNodeRequest()
	if err != nil {
		log.Println("Error at The DeleteNodeController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeleteNodeEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The DeleteNodeController")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Done The DeleteNodeController in RE")
	return
}

func CreateClusterController(rw http.ResponseWriter, req *http.Request) {
	log.Println("Inside The CreateClusterController in RE")
	eventsRequest := models.CreateCluster{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateCreateCluster(req.Body)
	if err != nil {
		log.Println("Error at The CreateClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateCreateClusterRequest()
	if err != nil {
		log.Println("Error at The CreateClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.CreateClusterEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The CreateClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Done The CreateClusterController in RE")
	return
}

func DeleteClusterController(rw http.ResponseWriter, req *http.Request){
	log.Println("Inside The DeleteClusterController in RE")
	eventsRequest := models.DeleteCluster{}
	successResponse := responses.Response{}
	err := eventsRequest.PopulateDeleteCluster(req.Body)
	if err != nil {
		log.Println("Error at The DeleteClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	err = eventsRequest.ValidateDeleteClusterRequest()
	if err != nil {
		log.Println("Error at The DeleteClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	ctx := context.Background()
	err = service.DeleteClusterEvents(ctx, &eventsRequest)
	if err != nil {
		log.Println("Error at The DeleteClusterController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, 200)
	log.Println("Delete The DeleteClusterController in RE")
	return
}

func ReloadRuleSetController(rw http.ResponseWriter, req *http.Request){
	log.Println("Inside The ReloadRuleSetController in RE")
	successResponse := responses.Response{}
    reloadRuleSet := models.ReloadRuleSet{}
	err := reloadRuleSet.PopulateRuleSet(req.Body)
	if err != nil{
		log.Println("Error at The ReloadRuleSetController in RE")
		log.Println("Error: ",err)
		responses.HandleError(rw, err)
		return
	}
	reloadRuleSet.ValidateAndResetPath()
	ctx := context.Background()
    service.ReloadRuleSet(ctx, &reloadRuleSet)
	successResponse.SendResponse(rw, 200)
	log.Println("Done The ReloadRuleSetController in RE")
	return
}





