package rc_etcd

import "go.etcd.io/etcd/contrib/recipes"

const SystemAppIngress = "system.ingress.app"

type IngressMessage struct {
	Ingress     []string
	ClusterName string
	SystemApps  []string
}

type ClusterConfig struct {
	ClusterName   string
	ClusterConfig string
}

func NewQueue(etcdClient *EtcdClient, prefix string) (*recipe.Queue, error) {
	queue := recipe.NewQueue(etcdClient.client, prefix)
	return queue, nil
}
