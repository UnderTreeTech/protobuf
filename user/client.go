package user

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

//SLA Definition
//API max handle duration: 500ms

// AppID Service.
const AppID = "service.user.v1"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (UserClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewUserClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
