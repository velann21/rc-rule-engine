package rc_etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/pkg/transport"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
	"go.etcd.io/etcd/clientv3"
)

var (
	ErrNotFound = errors.New("not found")
)

const (
	ClientPEM = "/mnt/etcd/ssl/certs/client.pem"
	ClientKey = "/mnt/etcd/ssl/certs/client-key.pem"
	TrustedCA = "/mnt/etcd/ssl/certs/ca.pem"
	CLOUD_INFO = "/cloudinfo"
)

type EtcdClient struct {
	client *clientv3.Client
}

type HostEntry struct {
	IP       string
	HostName string
}

func New(url []string) (*EtcdClient, error) {

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
		Endpoints:   url,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EtcdClient{client: c}, nil
}

func (e *EtcdClient) DeleteDNSEntry(ctx context.Context) error {
	_, err := e.client.Delete(ctx, "/host-entries")
	return err
}

func (e *EtcdClient) WriteRootKey(ctx context.Context, appName string, key []byte) error {
	_, err := e.client.Put(ctx, "/service/"+appName+"/root.key", string(key))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetRootKey(ctx context.Context, appName string) ([]byte, error) {

	resp, err := e.client.Get(ctx, "/service/"+appName+"/root.key")
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	return resp.Kvs[0].Value, nil
}

func (e *EtcdClient) WriteRootCert(ctx context.Context, appName string, key []byte) error {
	_, err := e.client.Put(ctx, "/service/"+appName+"/root.pem", string(key))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetRootCert(ctx context.Context, appName string) ([]byte, error) {

	resp, err := e.client.Get(ctx, "/service/"+appName+"/root.pem")
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	return resp.Kvs[0].Value, nil
}

func (e *EtcdClient) WriteLeafCert(ctx context.Context, appName string, key []byte) error {
	_, err := e.client.Put(ctx, "/service/"+appName+"/leaf.key", string(key))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetLeafCert(ctx context.Context, appName string) ([]byte, error) {

	resp, err := e.client.Get(ctx, "/service/"+appName+"/leaf.pem")
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	return resp.Kvs[0].Value, nil
}

func (e *EtcdClient) WriteLeafKey(ctx context.Context, appName string, key []byte) error {
	_, err := e.client.Put(ctx, "/service/"+appName+"/leaf.key", string(key))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetLeafKey(ctx context.Context, appName string) ([]byte, error) {

	resp, err := e.client.Get(ctx, "/service/"+appName+"/leaf.key")
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	return resp.Kvs[0].Value, nil
}

func (e *EtcdClient) SaveClusterPublicKey(ctx context.Context, clusterName string, publicKey []byte) error {

	_, err := e.client.Put(ctx, "/cluster/%v/.secret/publicKey", string(publicKey))
	if err != nil {
		return err
	}

	return nil
}

type Node struct {
	NodeName string
	NodeIP   string
}

func (e *EtcdClient) GetNodes(ctx context.Context) ([]Node, error) {

	resp, err := e.client.Get(ctx, "/nodeInfo", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var nodes []Node

	if len(resp.Kvs) == 0 {
		return nil, errors.New("node not found")
	}

	for _, val := range resp.Kvs {
		nodes = append(nodes, Node{
			NodeName: strings.Replace(string(val.Key), "/nodeInfo/", "", -1),
			NodeIP:   string(val.Value),
		})
	}

	return nodes, nil
}

func (e *EtcdClient) GetNodeInfo(ctx context.Context, nodeName string) (string, error) {

	resp, err := e.client.Get(ctx, fmt.Sprintf("/nodeInfo/%v", nodeName))
	if err != nil {
		return "", err
	}

	if len(resp.Kvs) == 0 {
		return "", errors.New("node not found")
	}

	return string(resp.Kvs[0].Value), nil
}

func (e *EtcdClient) UpdateNodeInfo(ctx context.Context, nodeName string, nodeIP string) error {

	_, err := e.client.Put(ctx, fmt.Sprintf("/nodeInfo/%v", nodeName), nodeIP)
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) SaveClusterPrivateKey(ctx context.Context, clusterName string, privateKey []byte) error {

	_, err := e.client.Put(ctx, "/cluster/%v/.secret/privateKey", string(privateKey))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetClusterPublicKey(ctx context.Context, clusterName string) ([]byte, error) {
	resp, err := e.client.Get(ctx, "/cluster/%v/rcssl/publicKey")
	if err != nil {
		return nil, err
	}

	if resp.Count > 0 {
		if len(resp.Kvs) > 0 {
			return resp.Kvs[0].Value, nil
		}
	}

	return nil, errors.New("value not found")
}

func (e *EtcdClient) GetClusterPrivateKey(ctx context.Context, clusterName string) ([]byte, error) {
	resp, err := e.client.Get(ctx, "/cluster/%v/rcssl/privateKey")
	if err != nil {
		return nil, err
	}

	if resp.Count > 0 {
		if len(resp.Kvs) > 0 {
			return resp.Kvs[0].Value, nil
		}
	}

	return nil, errors.New("value not found")
}

func (e *EtcdClient) GetCloudInfo(ctx context.Context) ([]byte, error) {
	resp, err := e.client.Get(ctx, CLOUD_INFO)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, errors.New("error while fetching cloud info")
	}

	return resp.Kvs[0].Value, nil
}

func dedup(hostEntries []HostEntry) []HostEntry {

	var newHostEntries []HostEntry

	for _, h := range hostEntries {
		for i, newH := range newHostEntries {
			if newH.HostName == h.HostName {
				newHostEntries = append(newHostEntries[:i], newHostEntries[i+1:]...)
			}
		}

		newHostEntries = append(newHostEntries, h)
	}

	return newHostEntries
}

func contains(a string, services []string) bool {

	for _, service := range services {
		if service == a {
			return true
		}
	}

	return false
}

func notApplicableForRemoval(hostEntries []HostEntry, serviceNames []string) []HostEntry {

	var out []HostEntry
	for _, entry := range hostEntries {
		if !contains(entry.HostName, serviceNames) {
			out = append(out, entry)
		}
	}

	return out
}

func notIn(serviceNames []string, hostname string) bool {
	for _, service := range serviceNames {
		if hostname == service {
			return false
		}
	}

	return true
}

func (e *EtcdClient) RemoveDNSEntry(ctx context.Context, serviceNames ...string) error {

	hostEntries, err := e.GetAllHostEntries(ctx)
	if err != nil {
		return err
	}

	var leftout []HostEntry

	for _, hostEntry := range hostEntries {
		if notIn(serviceNames, hostEntry.HostName) {
			leftout = append(leftout, hostEntry)
		}
	}

	//finalEntries := dedup(append(hostEntries, notApplicableForRemoval(hostEntries, serviceNames)...))

	entries, err := json.Marshal(&leftout)
	if err != nil {
		return err
	}

	_, err = e.client.Put(ctx, "/host-entries", string(entries))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) AddDNSEntry(ctx context.Context, entry ...HostEntry) error {

	hostEntries, err := e.GetAllHostEntries(ctx)
	if err != nil {
		return err
	}

	finalEntries := dedup(append(hostEntries, entry...))

	entries, err := json.Marshal(&finalEntries)
	if err != nil {
		return err
	}

	_, err = e.client.Put(ctx, "/host-entries", string(entries))
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdClient) GetAllHostEntries(ctx context.Context) ([]HostEntry, error) {
	var allHostEntries []HostEntry

	resp, err := e.client.Get(ctx, "/host-entries")
	if err != nil {
		return allHostEntries, err
	}

	if resp.Count != 0 {
		if len(resp.Kvs) > 0 {
			err = json.Unmarshal(resp.Kvs[0].Value, &allHostEntries)
			if err != nil {
				return allHostEntries, err
			}
		} else {
			return allHostEntries, errors.New("no key value found")
		}
	}

	return allHostEntries, nil
}

func (e *EtcdClient) FetchKeys() (*api.InitResponse, error) {

	var secret api.InitResponse

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := e.client.Get(ctx, "/rc-token")
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
		return nil, errors.New("node is null")
	}

	err = json.Unmarshal([]byte(resp.Kvs[0].Value), &secret)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func (e *EtcdClient) AddKeys(secret api.InitResponse) error {

	data, err := json.Marshal(&secret)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	_, err = e.client.Put(ctx, "/rc-token", string(data))
	cancel()

	if err != nil {
		return err
	}

	return nil
}
