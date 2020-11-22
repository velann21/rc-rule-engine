package helpers

import (
	"os"
	"testing"
)

func TestPromethuesMail_Alert(t *testing.T) {
	os.Setenv("ALERT_MANAGER_ENDPOINTS","http://localhost:9093")
	alertManagerFactoryProd := AlertManagerFactoryProducer{}
	alertManagerFactoryProd.GetAlertManagerFactory(PrometheusAMC).GetAlertType("Mail").Alert(map[string]string{
		"alertname":"testalert","EventsOccured":"sdsd","ErrorCode":"500","TraceID":"sdsdsdsdsddsd","EventType":"AddNode","ServiceName":"DC", "severity":"warning"})
}

