package rc_etcd

import (
	"context"
	"errors"
	"fmt"
)

var (
	CLUSTER_CONFIG   = "/cluster/%v/config"
	errClusterConfig = func(clusterName string) error {
		return errors.New("config not found for " + clusterName)
	}
)

func (e *EtcdClient) GetKubeConfig(ctx context.Context, clusterName string) ([]byte, error) {

	resp, err := e.client.Get(ctx, fmt.Sprintf(CLUSTER_CONFIG, clusterName))
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, errClusterConfig(clusterName)
	}

	if resp.Kvs[0] == nil {
		return nil, errClusterConfig(clusterName)
	}

	return resp.Kvs[0].Value, nil
}
