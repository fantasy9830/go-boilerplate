package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/fantasy9830/go-boilerplate/protos/helloworld"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// GrpcController ...
type GrpcController struct{}

const (
	address = "localhost:50051"
)

// SayHello ...
func (ctrl *GrpcController) SayHello(c *gin.Context) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "your name"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.String(http.StatusOK, res.Message)
}
