package usecases

import (
	"github.com/thelemmon/jokes/entities"
	"github.com/thelemmon/jokes/ports"
)

const JOKESBUFFERSIZE = 5 // represents the buffer size for fetching jokes as a margin of error to avoid duplicates

type GetUniqueRandomJokesUseCase struct {
	port ports.JokePort
}

func NewGetUniqueRandomJokesUseCase(port ports.JokePort) *GetUniqueRandomJokesUseCase {
	return &GetUniqueRandomJokesUseCase{
		port: port,
	}
}

func (useCase *GetUniqueRandomJokesUseCase) Execute(batchSize int) ([]*entities.Joke, error) {
	result := make([]*entities.Joke, 0, batchSize)
	// process jokes to remove duplicates
	for len(result) < batchSize {
		jokes, err := useCase.port.GetJokes(batchSize + JOKESBUFFERSIZE)
		if err != nil {
			return nil, err
		}
		for _, joke := range jokes {
			// if the result slice is full, break out of the loop
			if len(result) >= batchSize {
				break
			}
			// if the joke is already in the result slice, skip it
			if contains(result, joke) {
				continue
			}
			// if the joke is not already in the result slice, append it
			result = append(result, joke)
		}
	}
	return result, nil
}

func contains(slice []*entities.Joke, joke *entities.Joke) bool {
	for _, item := range slice {
		if item.ID == joke.ID {
			return true
		}
	}
	return false
}
