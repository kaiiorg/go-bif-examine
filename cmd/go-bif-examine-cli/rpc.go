package main

import (
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type RpcClient struct {
	grpcClient pb.BifExamineClient
}

func NewRpcClient(server string) (*RpcClient, error) {
	conn, err := grpc.Dial(
		server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	rpcClient := &RpcClient{
		grpcClient: pb.NewBifExamineClient(conn),
	}
	return rpcClient, nil
}
