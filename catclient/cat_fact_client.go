package catclient

import (
	"encoding/json"
	"io"
	"net/http"
)

const catFactURL string = "https://meowfacts.herokuapp.com/"

type CatFactClient interface {
	GetRandomCatFact(factChannel chan<- string)
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

func (catFactClientImpl *catFactClientImpl) GetRandomCatFact(factChannel chan<- string) {
	res, err := http.Get(catFactClientImpl.catFactsApiUrl)
	if err != nil { // or status code not 200
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	catFact := catFactResponse{}
	json.Unmarshal(body, &catFact)
	factChannel <- catFact.Data[0]
}
