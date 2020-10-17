package client

import (
	"fmt"
	"testing"

	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/example/V1/vo"
)

func TestClient_Call(t *testing.T) {
	discovery := discovery.NewP2PDiscovery("127.0.0.1:9000")

	client, err := NewClient(discovery)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := vo.AddUserReq{Name: "aaa"}
	res := &vo.AddUserRes{}

	err = client.Call("user.Add", req, res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
