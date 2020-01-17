package routes

import (
	"github.com/gorilla/mux"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/controller"
)

func Intialize(indexRoute *mux.Router) {
	indexRoute.HandleFunc("/syncApps_rules", controller.SyncAppsController).Methods("POST")
	indexRoute.HandleFunc("/deployApps_rules", controller.DeployAppsController).Methods("POST")
	indexRoute.HandleFunc("/addNode_rules", controller.AddNodeController).Methods("POST")
	indexRoute.HandleFunc("/deleteNode_rules", controller.DeleteNodeController).Methods("POST")
	indexRoute.HandleFunc("/createCluster_rules", controller.CreateClusterController).Methods("POST")
	indexRoute.HandleFunc("/deleteCluster_rules", controller.DeleteClusterController).Methods("POST")
	indexRoute.HandleFunc("/reload_rules", controller.ReloadRuleSetController).Methods("PUT")
}

