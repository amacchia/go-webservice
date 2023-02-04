package service

import (
	"testing"

	"example.com/webservice/model"
)

type MockDogFactService struct{}
type MockCatFactService struct{}

const expectedDogFact = "Some random dog fact"
const expectedCatFact = "Some random cat fact"

func (mockDogFactService *MockDogFactService) GetRandomDogFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{AnimalFact: expectedDogFact}
}

func (mockCatFactService *MockCatFactService) GetRandomCatFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{AnimalFact: expectedCatFact}
}

func TestRetrieveAnimalFacts(t *testing.T) {
	retrieveAnimalFactsService := NewAnimalFactsService()
	retrieveAnimalFactsService.dogFactClient = &MockDogFactService{}
	retrieveAnimalFactsService.catFactClient = &MockCatFactService{}

	actualAnimalFacts := retrieveAnimalFactsService.RetrieveAnimalFacts()
	actualDogFact := actualAnimalFacts.DogFact
	actualCatFact := actualAnimalFacts.CatFact

	if actualDogFact != "Some random dog fact" {
		t.Errorf("Expected dog fact to be Some random dog fact, but was %s", actualDogFact)
	}
	if actualCatFact != "Some random cat fact" {
		t.Errorf("Expected cat fact to be Some random cat fact, but was %s", actualCatFact)
	}
}
