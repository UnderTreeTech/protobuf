package stardust

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

//SLA Definition
//API max handle duration: 10ms

// AppID Service.
const AppID = "service.stardust.v1"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (StarDustClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewStarDustClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
