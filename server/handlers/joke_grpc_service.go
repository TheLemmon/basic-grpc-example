package handlers

import (
	"context"

	"github.com/thelemmon/jokes/jokes"
	usecases "github.com/thelemmon/jokes/use_cases"
)

type JokeGrpcService struct {
	jokes.UnimplementedJokeServiceServer
	useCase *usecases.GetUniqueRandomJokesUseCase
}

func NewJokeGrpcService(useCase *usecases.GetUniqueRandomJokesUseCase) *JokeGrpcService {
	return &JokeGrpcService{useCase: useCase}
}

func (service *JokeGrpcService) GetJokes(ctx context.Context, req *jokes.GetJokesRequest) (*jokes.GetJokesResponse, error) {
	result, err := service.useCase.Execute(int(req.BatchSize))
	if err != nil {
		return nil, err
	}
	// convert jokes entities to grpc response
	response := []*jokes.JokeResponse{}
	for _, joke := range result {
		response = append(response, &jokes.JokeResponse{
			Id:    joke.ID,
			Url:   joke.Url,
			Value: joke.Content,
		})
	}
	return &jokes.GetJokesResponse{Jokes: response}, nil
}
