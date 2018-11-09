package etcdcli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

func NewConnV3(data []string) (*etcdCli, error) {
	var err error
	conn := &etcdCli{}
	conn.CliV3, err = clientv3.New(clientv3.Config{
		Endpoints:   data,
		DialTimeout: DialTimeout,
	})
	return conn, err
}

func NewConnAuthV3(data []string, username, password string) (*etcdCli, error) {
	var err error
	conn := &etcdCli{}
	conn.CliV3, err = clientv3.New(clientv3.Config{
		Endpoints:   data,
		DialTimeout: DialTimeout,
		Username:    username,
		Password:    password,
	})
	return conn, err
}

func (this *etcdCli) PutV3(key, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Put(ctx, key, value)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatal(fmt.Printf("ctx is canceled by another routine: %v\n", err))
		case context.DeadlineExceeded:
			log.Fatal(fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err))
		case rpctypes.ErrEmptyKey:
			log.Fatal(fmt.Printf("client-side error: %v\n", err))
		default:
			log.Fatal(fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err))
		}
	}
	return resp, err
}

func (this *etcdCli) PutTtlV3(key, value string, ttl int64) (*clientv3.PutResponse, error) {
	TtlValue, err := this.CliV3.Grant(context.TODO(), ttl)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Put(ctx, key, value, clientv3.WithLease(TtlValue.ID))
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatal(fmt.Printf("ctx is canceled by another routine: %v\n", err))
		case context.DeadlineExceeded:
			log.Fatal(fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err))
		case rpctypes.ErrEmptyKey:
			log.Fatal(fmt.Printf("client-side error: %v\n", err))
		default:
			log.Fatal(fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err))
		}
	}
	return resp, err
}
func (this *etcdCli) GetV3(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *etcdCli) GetWithPreifxV3(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *etcdCli) DeleteV3(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Delete(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *etcdCli) DeleteWithPrefixV3(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.CliV3.Delete(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}
