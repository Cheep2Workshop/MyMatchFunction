package mmf

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"
)

type MatchFunctionService struct {
	grpc               *grpc.Server
	queryServiceClient pb.QueryServiceClient
	port               int
}

func Start(queryServiceAddr string, serverPort int) {
	conn, err := grpc.Dial(queryServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Open Match, got %s", err.Error())
	}
	defer conn.Close()

	mmfService := MatchFunctionService{
		queryServiceClient: pb.NewQueryServiceClient(conn),
	}

	server := grpc.NewServer()
	pb.RegisterMatchFunctionServer(server, &mmfService)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("TCP net listener initialization failed for port %v, got %s", serverPort, err.Error())
	}

	log.Printf("TCP net listener initialized for %v", serverPort)
	err = server.Serve(ln)
	if err != nil {
		log.Fatalf("gRPC serve failed, got %s", err.Error())
	}
}
