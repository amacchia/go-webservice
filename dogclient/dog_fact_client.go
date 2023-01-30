package dogclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"example.com/webservice/model"
)

const dogFactURL string = "https://dog-api.kinduff.com"

type DogFactClient interface {
	GetRandomDogFact(factChannel chan<- model.AnimalFactResult)
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

func (dogFactClientImpl *dogFactClientImpl) GetRandomDogFact(factChannel chan<- model.AnimalFactResult) {
	url := fmt.Sprintf("%s/api/facts", dogFactClientImpl.dogFactsApiServerUrl)
	animalFactResult := model.AnimalFactResult{}

	res, err := http.Get(url)
	if err != nil {
		animalFactResult.Error = errors.New("error retrieving dog fact")
		factChannel <- animalFactResult
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		animalFactResult.Error = fmt.Errorf("%d status code returned from dog fact API", res.StatusCode)
		factChannel <- animalFactResult
		return
	}

	body, _ := io.ReadAll(res.Body)
	dogFact := dogFactResponse{}
	err = json.Unmarshal(body, &dogFact)
	if err != nil {
		animalFactResult.Error = errors.New("error parsing dog fact response body")
		factChannel <- animalFactResult
		return
	}
	animalFactResult.AnimalFact = dogFact.Facts[0]
	factChannel <- animalFactResult
}
