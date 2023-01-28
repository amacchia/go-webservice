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
		catFactClient: catclient.NewCatClient(),
	}
}

func (animalFactsServiceImpl *animalFactsServiceImpl) RetrieveAnimalFacts() *model.AnimalFacts {
	dogFactChannel := make(chan string)
	catFactChannel := make(chan string)
	numberOfChannels := 2

	go animalFactsServiceImpl.dogFactClient.GetRandomDogFact(dogFactChannel)
	go animalFactsServiceImpl.catFactClient.GetRandomCatFact(catFactChannel)

	return collectAnimalFacts(numberOfChannels, dogFactChannel, catFactChannel)
}

func collectAnimalFacts(numberOfChannels int, dogFactChannel <-chan string, catFactChannel <-chan string) *model.AnimalFacts {
	var dogFact, catFact string

	for i := 0; i < numberOfChannels; i++ {
		select {
		case dogFact = <-dogFactChannel:
		case catFact = <-catFactChannel:
		}
	}

	return &model.AnimalFacts{DogFact: dogFact, CatFact: catFact}
}
