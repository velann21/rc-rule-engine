package helpers

import (
	"context"
	"fmt"
	"github.com/myntra/roulette/log"
	"gitlab.reynencourt.com/reynen-court/rc-common-lib/prometheus"
	"os"
	"time"
)

const (
	PrometheusAMC = "PrometheusAMC"
)
type AlertType interface {
	Alert(metaInformation map[string]string)
}

type PromethuesMail struct {

}

func (mail *PromethuesMail) Alert(metaInformation map[string]string) {
	go func(){
		p := prometheus.New(os.Getenv("ALERT_MANAGER_ENDPOINTS"))
		fmt.Println("Alert Sent")
		err := p.AddAlerts(context.Background(), &prometheus.Alert{
			Label: metaInformation,
			Annotation: map[string]string{"message":"test"},
			StartedAt:time.Now().UTC().Format(time.RFC3339Nano),
			EndsAt:time.Now().Add(10*time.Minute).UTC().Format(time.RFC3339Nano),
			GeneratorURL:"test",
		})
		if err != nil {
			log.Error("Something wrong while sending an alert")
		}
	}()

}

type OtherMail struct {

}

func (mail *OtherMail) Alert(metaInformation map[string]string) {

}

type AbstractAlertTypeFactory interface {
	GetAlertType(alertType string) AlertType
}

type PrometheusAlertTypeFactory struct {
    AlertType string
}

func (alertTypeFactory *PrometheusAlertTypeFactory) GetAlertType(alertType string) (AlertType){
    if alertType == "Mail"{
        return &PromethuesMail{}
	}
    return nil
}

type OtherAlertTypeFactory struct {
	AlertType string
}

func (alertTypeFactory *OtherAlertTypeFactory) GetAlertType(alertType string) (AlertType){
	if alertType == "Mail"{
		return &PromethuesMail{}
	}
	return nil
}


type AlertManagerFactoryProducer struct {
	FactoryType string
}

func (alertManagerFactoryProducer *AlertManagerFactoryProducer) GetAlertManagerFactory(factoryType string) AbstractAlertTypeFactory {
	if factoryType == PrometheusAMC{
		return &PrometheusAlertTypeFactory{}
	}
	return nil
}
