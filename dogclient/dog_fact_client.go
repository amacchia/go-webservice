package dogclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const dogFactURL string = "https://dog-api.kinduff.com"

type DogFactClient interface {
	GetRandomDogFact(factChannel chan<- string)
}

type dogFactClientImpl struct {
	dogFactsApiServerUrl string
}

type dogFactResponse struct {
	Facts   []string
	Success bool
}

func NewDogFactClient() *dogFactClientImpl {
	return &dogFactClientImpl{dogFactURL}
}

func (dogFactClientImpl *dogFactClientImpl) GetRandomDogFact(factChannel chan<- string) {
	url := fmt.Sprintf("%s/api/facts", dogFactClientImpl.dogFactsApiServerUrl)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	dogFact := dogFactResponse{}
	json.Unmarshal(body, &dogFact)
	factChannel <- dogFact.Facts[0]
}
