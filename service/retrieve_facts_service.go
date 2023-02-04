package service

import (
	"example.com/webservice/catclient"
	"example.com/webservice/dogclient"
	"example.com/webservice/model"
)

type AnimalFactsService interface {
	RetrieveAnimalFacts() *model.AnimalFacts
}

type animalFactsServiceImpl struct {
	dogFactClient dogclient.DogFactClient
	catFactClient catclient.CatFactClient
}

func NewAnimalFactsService() *animalFactsServiceImpl {
	return &animalFactsServiceImpl{
		dogFactClient: dogclient.NewDogFactClient(),
		catFactClient: catclient.NewCatFactClient(),
	}
}

func (animalFactsServiceImpl *animalFactsServiceImpl) RetrieveAnimalFacts() *model.AnimalFacts {
	dogFactChannel := make(chan model.AnimalFactResult)
	catFactChannel := make(chan model.AnimalFactResult)
	numberOfChannels := 2

	go animalFactsServiceImpl.dogFactClient.GetRandomDogFact(dogFactChannel)
	go animalFactsServiceImpl.catFactClient.GetRandomCatFact(catFactChannel)

	return collectAnimalFacts(numberOfChannels, dogFactChannel, catFactChannel)
}

func collectAnimalFacts(numberOfChannels int, dogFactChannel <-chan model.AnimalFactResult, catFactChannel <-chan model.AnimalFactResult) *model.AnimalFacts {
	var dogFactResult, catFactResult model.AnimalFactResult

	for i := 0; i < numberOfChannels; i++ {
		select {
		case dogFactResult = <-dogFactChannel:
		case catFactResult = <-catFactChannel:
		}
	}

	return &model.AnimalFacts{DogFact: dogFactResult.AnimalFact, CatFact: catFactResult.AnimalFact}
}
