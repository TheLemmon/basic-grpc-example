package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/thelemmon/jokes/entities"
)

type JokeHttpAdapter struct {
	url string
}

func NewJokeHttpAdapter() *JokeHttpAdapter {
	return &JokeHttpAdapter{
		url: "https://api.chucknorris.io/jokes/random",
	}
}

func (adapter *JokeHttpAdapter) GetJokes(batchSize int) ([]*entities.Joke, error) {
	channel := make(chan *entities.Joke)
	result := make([]*entities.Joke, 0, batchSize)

	// Start a goroutine to fetch jokes concurrently
	for i := 0; i < batchSize; i++ {
		go adapter.getJoke(channel)
	}
	for i := 0; i < batchSize; i++ {
		joke := <-channel
		result = append(result, joke)
	}
	return result, nil
}

func (adapter *JokeHttpAdapter) getJoke(channel chan<- *entities.Joke) {
	joke := &entities.Joke{}
	response, err := http.Get(adapter.url)
	if err != nil {
		channel <- joke
		return
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(joke)
	channel <- joke
}
