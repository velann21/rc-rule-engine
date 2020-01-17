package rc_etcd

import (
	"context"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

const (
	CLUSTERS        = "/clusters"
	DOWNLOADED_APPS = "/apps"
	DEPLOYED_APPS   = "/deployments"
)

type Cluster struct {
	ClusterName   string
	ClusterConfig string
	Ingress       string
}

/**

/deployments/clusterName/apps/deploymentName

*/
func (e *EtcdClient) GetAllDeployedApps(ctx context.Context) ([]string, error) {
	var resp []string

	res, err := e.client.Get(ctx, DEPLOYED_APPS, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return []string{}, nil
	}

	for _, v := range res.Kvs {

		keyNames := strings.Split(string(v.Key), "/")

		if len(keyNames) == 4 {

			resp = append(resp, keyNames[3])
		}
	}

	return resp, nil
}

func (e *EtcdClient) GetAllDownloadedApps(ctx context.Context) ([]string, error) {

	var apps = make(map[string]bool, 0)
	var resp []string

	res, err := e.client.Get(ctx, DOWNLOADED_APPS, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return []string{}, nil
	}

	for _, v := range res.Kvs {

		keyNames := strings.Split(string(v.Key), "/")

		if len(keyNames) > 1 {
			apps[keyNames[1]] = true
		}
	}

	for k := range apps {
		resp = append(resp, k)
	}

	return resp, nil
}

func (e *EtcdClient) GetAllClustersKey(ctx context.Context) ([]string, error) {

	var clusters = make(map[string]bool, 0)
	var resp []string

	res, err := e.client.Get(ctx, CLUSTERS, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return []string{}, nil
	}

	for _, v := range res.Kvs {

		keyNames := strings.Split(string(v.Key), "/")

		if len(keyNames) > 1 {
			clusters[keyNames[1]] = true
		}
	}

	for k := range clusters {
		resp = append(resp, k)
	}

	return resp, nil
}
