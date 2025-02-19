package ports

import "github.com/thelemmon/jokes/entities"

type JokePort interface {
	GetJokes(batchSize int) ([]*entities.Joke, error)
}
