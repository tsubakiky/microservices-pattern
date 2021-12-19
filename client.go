package main

import (
	"log"

	authority "github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := authority.NewAuthorityServiceClient(conn)

	ctx := context.Background()
	response, err := c.Signup(ctx, &authority.SignupRequest{Name: "goldie"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.GetCustomer().Name)
}
