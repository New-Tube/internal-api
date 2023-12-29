package main

import (
	"fmt"
	"internal-api/db"
	"internal-api/endpoints"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.ConnectToDB()
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatalf("Cannot migrate the database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 5050))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	endpoints.RegisterUserService(s)
	endpoints.RegisterCommentService(s)
	endpoints.RegisterVideoRegularUserService(s)
	endpoints.RegisterVideoCreatorUserService(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
