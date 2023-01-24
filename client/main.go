package main

import (
	"context"
	"log"
	pb "sample/proto/greetings"

	inter "github.com/vijeyash1/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	checkPolicy := &inter.CheckPolicy{}
	actions := []string{"greet"}
	checkPolicy = checkPolicy.AddActions(actions).AddScope("airtel").AddRole("admin").AddKind("greetings")
	log.Println("kind:", checkPolicy.GetKind())
	log.Println("scope", checkPolicy.GetScope())
	log.Println("role", checkPolicy.GetRole())
	log.Println("action", checkPolicy.GetAction())
	valid, err := checkPolicy.Valid()
	if err != nil {
		log.Fatalf("not a valid policy: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(valid.UnaryClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetingsServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "vijeyash"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
