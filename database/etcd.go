package database

import (
	"github.com/coreos/etcd/clientv3"
	"os"
	"strings"
	"time"
	"go.etcd.io/etcd/pkg/transport"

)
const (
	ClientPEM = "/mnt/etcd/ssl/certs/client.pem"
	ClientKey = "/mnt/etcd/ssl/certs/client-key.pem"
	TrustedCA = "/mnt/etcd/ssl/certs/ca.pem"
)
var etcdConnection  *clientv3.Client
func ConnectEtcd() (*clientv3.Client, error){

	etcdUrl := strings.Split(os.Getenv("ETCD_ADDR"), ",")
	tlsInfo := transport.TLSInfo{
		CertFile:      ClientPEM,
		KeyFile:       ClientKey,
		TrustedCAFile: TrustedCA,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, err
	}

	cfg := clientv3.Config{
		Endpoints:   etcdUrl,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
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


