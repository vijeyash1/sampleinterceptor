package main

import (
	"net"
	pb "sample/proto/greetings"

	inter "github.com/vijeyash1/interceptor"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetingsServiceServer
}

func main() {

	AddPolicyfirst()
	AddPolicysecond()
	cerbosconfig := inter.NewCerbosConfig("localhost:3593")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := Server{}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(cerbosconfig.UnaryServerInterceptor()))
	pb.RegisterGreetingsServiceServer(grpcServer, &s)
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
