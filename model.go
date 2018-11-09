package etcdcli

import (
	"time"

	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
)

const DialTimeout time.Duration = 5 * time.Second

type etcdCli struct {
	CliV3 *clientv3.Client
	CliV2 *client.Client
}
