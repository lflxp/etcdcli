package etcdcli

import (
	"testing"
)

func TestNewConnV3(t *testing.T) {
	data := []string{"http://localhost:2379"}
	_, err := NewConnV3(data, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

// func TestGetWithPreifxV2(t *testing.T) {
// 	data := []string{"http://localhost:2379"}
// 	cli, err := NewConnV2(data, "", "")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resp, err := cli.GetWithPreifxV2("/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(resp)
// }
