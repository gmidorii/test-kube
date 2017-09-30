package main

import (
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/midorigreen/test-kube/protoc"
	"golang.org/x/net/context"
)

type PingServerService struct{}

func (p PingServerService) Ok(ctx context.Context, r *pb.OkRequest) (*pb.OkResponse, error) {
	log.Println(r.Quetion)
	if r.Quetion != "ping" {
		return &pb.OkResponse{}, errors.New("not ping")
	}
	return &pb.OkResponse{Answer: "pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	serv := grpc.NewServer()
	pb.RegisterPingServer(serv, PingServerService{})
	serv.Serve(lis)
}
