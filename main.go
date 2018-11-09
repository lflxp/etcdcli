package etcdcli

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/coreos/etcd/clientv3"
// 	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
// 	"golang.org/x/net/context"
// )

// func main() {
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"localhost:2379"},
// 		DialTimeout: 5 * time.Second,
// 	})
// 	if err != nil {
// 		// handle error!
// 		fmt.Println(err.Error())
// 	}
// 	defer cli.Close()

// 	// resp, err := cli.MemberList(context.Background())
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println("members:", len(resp.Members), resp.Members[0].GetName())

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// _, err = cli.Put(ctx, time.Now().String(), "2")
// 	_, err = cli.Put(ctx, "/1/3", "2")
// 	if err != nil {
// 		switch err {
// 		case context.Canceled:
// 			fmt.Printf("ctx is canceled by another routine: %v\n", err)
// 		case context.DeadlineExceeded:
// 			fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
// 		case rpctypes.ErrEmptyKey:
// 			fmt.Printf("client-side error: %v\n", err)
// 		default:
// 			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
// 		}
// 	}

// 	resp, err := cli.Get(ctx, "", clientv3.WithPrefix())
// 	cancel()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, x := range resp.Kvs {
// 		fmt.Println(string(x.Key), string(x.Value))
// 	}
// }
