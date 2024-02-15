package worker

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/alexohneander/flotte/services/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func StartWorker() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTasksClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	ctx := context.Background()

	for {
		r, err := c.GetTasks(ctx, &pb.GetTasksRequest{PluginId: "wordcount"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Get Task: { %s }", r)

		_, err = c.MarkTaskAsDone(ctx, &pb.MarkTaskAsDoneRequest{Id: r.Id, TaskType: "map"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(3 * time.Second)
	}
}
