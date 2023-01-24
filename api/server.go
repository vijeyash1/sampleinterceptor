package main

import (
	"context"
	"fmt"
	"log"

	pb "sample/proto/greetings"

	"github.com/cerbos/cerbos/client"
)

func AddPolicyfirst() {
	cli, err := client.NewAdminClientWithCredentials("localhost:3593", "cerbos", "randomHash", client.WithPlaintext())
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	policyset := client.PolicySet{}
	actions := []string{
		"greet",
	}
	resource := client.NewAllowResourceRule(actions...).WithRoles("admin")
	
	resourcePolicy := client.NewResourcePolicy("greetings", "default").AddResourceRules(resource)

	resourcePolicy.WithScope("")

	policySet := policyset.AddResourcePolicies(resourcePolicy)
	err = cli.AddOrUpdatePolicy(context.Background(), policySet)
	if err != nil {
		log.Fatalf("failed to add policy: %v", err)
	}
	log.Println("Policy added successfully")
}
func AddPolicysecond() {
	cli, err := client.NewAdminClientWithCredentials("localhost:3593", "cerbos", "randomHash", client.WithPlaintext())
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	policyset := client.PolicySet{}
	actions := []string{
		"greet",
	}
	resource := client.NewAllowResourceRule(actions...).WithRoles("admin")
	resourcePolicy := client.NewResourcePolicy("greetings", "default").AddResourceRules(resource)

	resourcePolicy.WithScope("airtel")

	policySet := policyset.AddResourcePolicies(resourcePolicy)
	err = cli.AddOrUpdatePolicy(context.Background(), policySet)
	if err != nil {
		log.Fatalf("failed to add policy: %v", err)
	}
	log.Println("Policy added successfully")
}


func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: fmt.Sprintf("hello , %s", in.Name)}, nil
}
