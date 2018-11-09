package etcdcli

import (
	"context"
	"time"

	"github.com/coreos/etcd/client"
)

func NewConnV2(data []string) (*etcdCli, error) {
	var err error
	conn := &etcdCli{}
	*conn.CliV2, err = client.New(client.Config{
		Endpoints:               data,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	})
	return conn, err
}

func NewConnAuthV2(data []string, username, password string) (*etcdCli, error) {
	var err error
	conn := &etcdCli{}
	*conn.CliV2, err = client.New(client.Config{
		Endpoints:               data,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
		Username:                username,
		Password:                password,
	})
	return conn, err
}

func (this *etcdCli) GetKeysApi() client.KeysAPI {
	return client.NewKeysAPI(*this.CliV2)
}

func (this *etcdCli) PutV2(key, value string) (*client.Response, error) {
	resp, err := this.GetKeysApi().Set(context.Background(), key, value, &client.SetOptions{})
	return resp, err
}

func (this *etcdCli) PutTtlV2(key, value string, ttl int64) (*client.Response, error) {
	resp, err := this.GetKeysApi().Set(context.Background(), key, value, &client.SetOptions{TTL: time.Duration(ttl) * time.Second})
	return resp, err
}
func (this *etcdCli) GetV2(key string) (*client.Response, error) {
	resp, err := this.GetKeysApi().Get(context.Background(), key, &client.GetOptions{})
	return resp, err
}

func (this *etcdCli) GetWithPreifxV2(key string) (*client.Response, error) {
	resp, err := this.GetKeysApi().Get(context.Background(), key, &client.GetOptions{Recursive: true})
	return resp, err
}

func (this *etcdCli) DeleteV2(key string) (*client.Response, error) {
	resp, err := this.GetKeysApi().Delete(context.Background(), key, &client.DeleteOptions{})
	return resp, err
}

func (this *etcdCli) DeleteWithPrefixV2(key string) (*client.Response, error) {
	resp, err := this.GetKeysApi().Delete(context.Background(), key, &client.DeleteOptions{Recursive: true})
	return resp, err
}
