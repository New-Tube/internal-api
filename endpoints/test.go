package endpoints

import (
	"context"
	pb "internal-api/protos/test"

	"google.golang.org/grpc"
)

type testService struct {
	pb.UnimplementedTestServer
}

func (s *testService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + req.Name + "!",
	}, nil
}

func RegisterTestService(s *grpc.Server) {
	pb.RegisterTestServer(s, &testService{})
}
