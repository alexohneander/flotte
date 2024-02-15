package coordinator

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/alexohneander/flotte/services/tasks"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type tasksServer struct {
	pb.UnimplementedTasksServer
}

func (s *tasksServer) GetTasks(ctx context.Context, request *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	log.Printf("Received Request: %v", request)
	resp := &pb.GetTasksResponse{
		Id:       uuid.New().String(),
		Name:     "grimm.txt",
		Number:   20,
		Type:     "wordcount",
		TaskType: "map",
		Nreduce:  2,
	}

	return resp, nil
}

func (s *tasksServer) MarkTaskAsDone(ctx context.Context, request *pb.MarkTaskAsDoneRequest) (*pb.MarkTaskAsDoneResponse, error) {
	log.Printf("Received Task completion: { %v }", request)
	resp := &pb.MarkTaskAsDoneResponse{
		Status: "OK",
	}

	return resp, nil
}

func StartCoordinator() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTasksServer(s, &tasksServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
