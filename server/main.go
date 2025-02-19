package main

import (
	"log"
	"net"

	"github.com/thelemmon/jokes/adapters"
	"github.com/thelemmon/jokes/handlers"
	"github.com/thelemmon/jokes/jokes"
	usecases "github.com/thelemmon/jokes/use_cases"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpServer := grpc.NewServer()

	// Dependency injection
	adapter := adapters.NewJokeHttpAdapter()
	useCase := usecases.NewGetUniqueRandomJokesUseCase(adapter)
	service := handlers.NewJokeGrpcService(useCase)

	// Register Grpc service
	jokes.RegisterJokeServiceServer(grpServer, service)

	// Start the server
	log.Printf("Server started on port %s", ":5050")
	if err := grpServer.Serve(listener); err != nil {
		panic(err)
	}

}
