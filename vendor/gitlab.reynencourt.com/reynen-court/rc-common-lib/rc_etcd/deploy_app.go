package rc_etcd

//import (
//	"encoding/json"
//
//	"gitlab.reynencourt.com/reynen-court/rc-common-lib/proto"
//	recipe "go.rc_etcd.io/etcd/contrib/recipes"
//)
//
//const DEPLOY_APP_QUEUE_PREFIX = "rc.deploy.queue"
//
//func (e *EtcdClient) RequestAppDeployment(deployment *proto.AppDeploymentRequest) error {
//
//	q, err := NewQueue(e, DEPLOY_APP_QUEUE_PREFIX)
//	if err != nil {
//		return err
//	}
//
//	d, err := json.Marshal(deployment)
//
//	if err != nil {
//		return err
//	}
//
//	return q.Enqueue(string(d))
//}
//
//func GetDeploymentRequest(q *recipe.Queue) (*proto.AppDeploymentRequest, error) {
//	msg, err := q.Dequeue()
//
//	if err != nil {
//		return nil, err
//	}
//
//	var deploymentReq proto.AppDeploymentRequest
//
//	err = json.Unmarshal([]byte(msg), &deploymentReq)
//
//	if err != nil {
//		return nil, err
//	}
//	return &deploymentReq, nil
//}
