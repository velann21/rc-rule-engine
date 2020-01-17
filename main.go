package main

import (
	"github.com/gorilla/mux"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/routes"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/service"
	//etcdConn "gitlab.reynencourt.com/reynen-court/rc-common-lib/rc_etcd"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/database"
	"log"
	"net/http"
	"os"
	//"strings"
)


func main() {
	os.Setenv("ETCD_ADDR", "http://localhost:2379, http://localhost:2379")
	os.Setenv("ALERT_MANAGER_ENDPOINTS","http://localhost:9093")
	os.Setenv("RULESET_FILEPATH","dsl/events.xml")
	r := mux.NewRouter().StrictSlash(false)
	mainRoutes := r.PathPrefix("/api/v1/rule_engine").Subrouter()
	ruleSet := helpers.RuleSet{}
	//etcdAddresses := strings.Split(os.Getenv("ETCD_ADDR"), ",")
	//_, err := etcdConn.New(etcdAddresses)
	_, err := database.ConnectEtcd()
	if err != nil{
		log.Println("Etcd may be down")
		os.Exit(3)
	}
	go service.ExecuteEventForNotification()
	ruleSet.LoadRuleSet(os.Getenv("RULESET_FILEPATH"))
	routes.Intialize(mainRoutes)
	log.Println("starting server..")
	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}
