package database

import (
	"github.com/coreos/etcd/clientv3"
	"os"
	"strings"
	"time"
)
var etcdConnection  *clientv3.Client
func ConnectEtcd() (*clientv3.Client, error){
	var EtcdAddr = os.Getenv("ETCD_ADDR")
	cfg := clientv3.Config{
		Endpoints:   strings.Split(EtcdAddr, ","),
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	etcdConnection = client
    return client, err
}

func GetEtcdConnection() *clientv3.Client{
     return etcdConnection
}


