package dao

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"gitlab.reynencourt.com/reynen-court/rc-rules-engine/helpers"
)

const EVENTS  = "/Events/"

func StoreEvent(ctx context.Context, client *clientv3.Client,eventType string, events []byte) error {
	_, err := client.Put(ctx, EVENTS+eventType, string(events))
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	return nil
}

func DeleteEvent(ctx context.Context, client *clientv3.Client,eventType string) error{
	delResp, err := client.Delete(ctx, EVENTS+eventType)
	if err != nil{
		return helpers.ErrSomeThingWentWrng
	}
	fmt.Println(delResp.Deleted)
	return nil
}

func GetEvent(ctx context.Context, client *clientv3.Client,eventType string) (*clientv3.GetResponse,error){
	resp, err := client.Get(ctx, EVENTS+eventType)
    if err != nil{
    	return nil, err
	}
	return resp, nil
}
