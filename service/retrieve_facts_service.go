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
	DogFactService dogclient.DogFactService
	CatFactService catclient.CatFactClient
}

func NewAnimalFactsService() *animalFactsServiceImpl {
	return &animalFactsServiceImpl{
		DogFactService: dogclient.NewDogFactClient(),
		CatFactService: catclient.NewCatClient(),
	}
}

func (animalFactsServiceImpl *animalFactsServiceImpl) RetrieveAnimalFacts() *model.AnimalFacts {
	dogFactChannel := make(chan string)
	catFactChannel := make(chan string)
	numberOfChannels := 2

	go animalFactsServiceImpl.DogFactService.GetRandomDogFact(dogFactChannel)
	go animalFactsServiceImpl.CatFactService.GetRandomCatFact(catFactChannel)

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
