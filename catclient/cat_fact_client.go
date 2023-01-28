package catclient

import (
	"encoding/json"
	"io"
	"net/http"

	"example.com/webservice/model"
)

const catFactURL string = "https://meowfacts.herokuapp.com/"

type CatFactClient interface {
	GetRandomCatFact(factChannel chan<- string)
}

type catFactClientImpl struct {
	catFactsApiUrl string
}

func NewCatClient() *catFactClientImpl {
	return &catFactClientImpl{catFactURL}

}

func (catFactClientImpl *catFactClientImpl) GetRandomCatFact(factChannel chan<- string) {
	res, err := http.Get(catFactClientImpl.catFactsApiUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Called after function runs

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	catFact := model.CatFactResponse{}
	json.Unmarshal(body, &catFact)
	factChannel <- catFact.Data[0]
}
