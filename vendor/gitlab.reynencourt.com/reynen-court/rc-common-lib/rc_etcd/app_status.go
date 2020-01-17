package rc_etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type DeployedAppMetadata struct {
	AppName        string    `json:"app_name"`
	DeploymentUUID string    `json:"deployment_uuid"`
	DeployedOn     time.Time `json:"deployed_on"`
	Status         string    `json:"status"`
	Version        string    `json:"version"`
	Cluster        string    `json:"cluster"`
}

type DeployedApp struct {
	Name      string              `json:"name"`
	Metatdata DeployedAppMetadata `json:"metatdata"`
}

func (e *EtcdClient) GetAllDeployedAppsInCluster(ctx context.Context, clusterName string) ([]DeployedApp, error) {

	var deployedApps []DeployedApp

	res, err := e.client.Get(ctx, fmt.Sprintf(DEPLOYED_APPS+"/%s/apps", clusterName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return nil, nil
	}

	for _, v := range res.Kvs {

		var app DeployedApp

		if err := json.Unmarshal(v.Value, &app); err != nil {
			return nil, err
		}

		deployedApps = append(deployedApps, app)
	}

	return deployedApps, nil
}

func (e *EtcdClient) SetStatusForDeployedApp(ctx context.Context, deploymentName string, clusterName string, status string) error {
	etcdKey := fmt.Sprintf(DEPLOYED_APPS+"/%s/apps/%s", clusterName, deploymentName)
	resp, err := e.client.Get(ctx, etcdKey)
	if err != nil {
		return err
	}

	if resp.Count == 0 {
		msg := fmt.Sprintf("app with deployment_name=%s not found in cluster=%s", deploymentName, clusterName)
		return errors.New(msg)
	}
	var app DeployedApp
	if err := json.Unmarshal(resp.Kvs[0].Value, &app); err != nil {
		return err
	}
	app.Metatdata.Status = status

	d, err := json.Marshal(&app)
	if err != nil {
		return err
	}

	_, err = e.client.Put(ctx, etcdKey, string(d))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetDeployedApp(ctx context.Context, clusterName, deploymentName string) (*DeployedApp, error) {

	var deployedApps DeployedApp

	res, err := e.client.Get(ctx, fmt.Sprintf(DEPLOYED_APPS+"/%s/apps/%s", clusterName, deploymentName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return nil, nil
	}

	var metadata DeployedAppMetadata
	keyNames := strings.Split(string(res.Kvs[0].Key), "/")
	if err := json.Unmarshal(res.Kvs[0].Value, &metadata); err != nil {
		return nil, err
	}
	deployedApps.Name = keyNames[3]
	deployedApps.Metatdata = metadata

	return &deployedApps, nil
}

func (e *EtcdClient) IsAppDeployed(ctx context.Context, clusterName, deploymentName string) (bool, error) {

	res, err := e.client.Get(ctx, fmt.Sprintf(DEPLOYED_APPS+"/%s/apps/%s", clusterName, deploymentName), clientv3.WithPrefix())
	if err != nil {
		return false, err
	}

	return res.Count > 0, nil

}

func (e *EtcdClient) SaveAppDetail(ctx context.Context, appDetail DeployedApp) error {

	d, err := json.Marshal(&appDetail)
	if err != nil {
		return err
	}

	_, err = e.client.Put(ctx, fmt.Sprintf(DEPLOYED_APPS+"/%s/apps/%s", appDetail.Metatdata.Cluster, appDetail.Name), string(d))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetDeployedAppDetailsAcrossClusters(ctx context.Context, deploymentName string) ([]DeployedApp, error) {

	var apps []DeployedApp

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := e.client.Get(ctx, DEPLOYED_APPS, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	}

	if resp.Kvs == nil {
		return nil, errors.New("app status is null")
	}

	for _, value := range resp.Kvs {
		keys := strings.Split(string(value.Key), "/")
		if keys[4] != deploymentName {
			continue
		}
		var app DeployedApp
		err := json.Unmarshal(value.Value, &app)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (e *EtcdClient) GetAllDeployedAppDetailsAcrossClusters(ctx context.Context) ([]DeployedApp, error) {

	var apps []DeployedApp

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := e.client.Get(ctx, DEPLOYED_APPS, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	}

	if resp.Kvs == nil {
		return nil, errors.New("app status is null")
	}

	for _, value := range resp.Kvs {
		var app DeployedApp
		err := json.Unmarshal(value.Value, &app)

		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (e *EtcdClient) RemoveDeployedApp(ctx context.Context, clusterName, deploymentName string) error {

	//TODO: I dont do any explicit check on the apps ATM
	if clusterName == "" || deploymentName == "" {
		return errors.New("cluster_name and deployment_name are required fields")
	}

	etcdKey := fmt.Sprintf(DEPLOYED_APPS+"/%s/apps/%s", clusterName, deploymentName)

	_, err := e.client.Delete(ctx, etcdKey)

	if err != nil {
		return err
	}
	return nil
}
