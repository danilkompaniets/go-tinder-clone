package utils

import (
	"github.com/danilkompanites/tinder-clone/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUsersClient() (gen.UserClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return gen.NewUserClient(conn), conn, nil
}
