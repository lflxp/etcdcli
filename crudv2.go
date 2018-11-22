package etcdcli

import (
	"context"
	"time"

	"github.com/etcd-io/etcd/client"
)

type V2 struct {
	Endpoints []string
	Username  string
	Password  string
	Cli       client.Client
	KeysApi   client.KeysAPI
}

// data  []string{"http://localhost:2379"}
func NewConnV2(data []string, username, password string) (*V2, error) {
	s := &V2{
		Endpoints: data,
		Username:  username,
		Password:  password,
	}
	err := s.Init()
	return s, err
}

func (this *V2) Init() error {
	cfg := client.Config{
		Endpoints:               this.Endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	}
	if this.Username != "" && this.Password != "" {
		cfg.Username = this.Username
		cfg.Password = this.Password
	}
	cli, err := client.New(cfg)
	if err != nil {
		return err
	}
	this.Cli = cli
	return nil
}

func (this *V2) getKeysApi() client.KeysAPI {
	if this.Cli == nil {
		this.Init()
	}
	if this.KeysApi == nil {
		this.KeysApi = client.NewKeysAPI(this.Cli)
	}
	return this.KeysApi
}

func (this *V2) PutV2(key, value string) (*client.Response, error) {
	resp, err := this.getKeysApi().Set(context.Background(), key, value, &client.SetOptions{})
	return resp, err
}

func (this *V2) PutTtlV2(key, value string, ttl int64) (*client.Response, error) {
	resp, err := this.getKeysApi().Set(context.Background(), key, value, &client.SetOptions{TTL: time.Duration(ttl) * time.Second})
	return resp, err
}
func (this *V2) GetV2(key string) (*client.Response, error) {
	resp, err := this.getKeysApi().Get(context.Background(), key, nil)
	return resp, err
}

func (this *V2) GetWithPreifxV2(key string) (*client.Response, error) {
	resp, err := this.getKeysApi().Get(context.Background(), key, &client.GetOptions{Recursive: true})
	return resp, err
}

func (this *V2) DeleteV2(key string) (*client.Response, error) {
	resp, err := this.getKeysApi().Delete(context.Background(), key, &client.DeleteOptions{})
	return resp, err
}

func (this *V2) DeleteWithPrefixV2(key string) (*client.Response, error) {
	resp, err := this.getKeysApi().Delete(context.Background(), key, &client.DeleteOptions{Recursive: true})
	return resp, err
}
