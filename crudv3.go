package etcdcli

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

type V3 struct {
	Endpoints []string
	Username  string
	Password  string
	Cli       *clientv3.Client
}

func NewConnV3(data []string, username, password string) (*V3, error) {
	conn := &V3{
		Endpoints: data,
		Username:  username,
		Password:  password,
	}
	err := conn.Init()
	return conn, err
}

func (this *V3) Init() error {
	cfg := clientv3.Config{
		Endpoints:   this.Endpoints,
		DialTimeout: 5 * time.Second,
	}
	// if this.Username != "" && this.Password != "" {
	// 	cfg.Username = this.Username
	// 	cfg.Password = this.Password
	// }
	cli, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	this.Cli = cli
	return nil
}

func (this *V3) PutV3(key, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Put(ctx, key, value)
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

func (this *V3) PutTtlV3(key, value string, ttl int64) (*clientv3.PutResponse, error) {
	TtlValue, err := this.Cli.Grant(context.TODO(), ttl)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Put(ctx, key, value, clientv3.WithLease(TtlValue.ID))
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
func (this *V3) GetV3(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *V3) GetWithPreifxV3(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *V3) DeleteV3(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Delete(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func (this *V3) DeleteWithPrefixV3(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := this.Cli.Delete(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}
