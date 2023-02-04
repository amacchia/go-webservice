package catclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"example.com/webservice/model"
)

const catFactURL string = "https://meowfacts.herokuapp.com/"

type CatFactClient interface {
	GetRandomCatFact(factChannel chan<- model.AnimalFactResult)
}

type catFactClientImpl struct {
	catFactsApiUrl string
}

type catFactResponse struct {
	Data []string
}

func NewCatFactClient() *catFactClientImpl {
	return &catFactClientImpl{catFactURL}
}

func (catFactClientImpl *catFactClientImpl) GetRandomCatFact(factChannel chan<- model.AnimalFactResult) {
	animalFactResult := model.AnimalFactResult{}

	res, err := http.Get(catFactClientImpl.catFactsApiUrl)
	if err != nil {
		animalFactResult.Error = errors.New("error retrieving cat fact")
		factChannel <- animalFactResult
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		animalFactResult.Error = fmt.Errorf("%d status code returned from cat fact API", res.StatusCode)
		factChannel <- animalFactResult
		return
	}

	body, _ := io.ReadAll(res.Body)
	catFact := catFactResponse{}
	err = json.Unmarshal(body, &catFact)
	if err != nil {
		animalFactResult.Error = errors.New("error parsing cat fact response body")
		factChannel <- animalFactResult
		return
	}
	animalFactResult.AnimalFact = catFact.Data[0]
	factChannel <- animalFactResult
}
