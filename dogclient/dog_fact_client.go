package dogclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/webservice/model"
)

const dogFactURL string = "https://dog-api.kinduff.com"

type DogFactService interface {
	GetRandomDogFact(factChannel chan<- string)
}

type dogFactClientImpl struct {
	dogFactsApiServerUrl string
}

func NewDogFactClient() *dogFactClientImpl {
	return &dogFactClientImpl{dogFactURL}
}

func (dogFactClientImpl *dogFactClientImpl) GetRandomDogFact(factChannel chan<- string) { // This fucntion can only send data into the fact channel
	url := fmt.Sprintf("%s/api/facts", dogFactClientImpl.dogFactsApiServerUrl)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Called after function runs

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	dogFact := model.DogFactResponse{}
	json.Unmarshal(body, &dogFact)
	factChannel <- dogFact.Facts[0]
}
