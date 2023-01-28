package service

import "testing"

type MockDogFactService struct{}
type MockCatFactService struct{}

const expectedDogFact = "Some random dog fact"
const expectedCatFact = "Some random cat fact"

func (mockDogFactService *MockDogFactService) GetRandomDogFact(factChannel chan<- string) {
	factChannel <- expectedDogFact
}

func (mockCatFactService *MockCatFactService) GetRandomCatFact(factChannel chan<- string) {
	factChannel <- expectedCatFact
}

func TestRetrieveAnimalFacts(t *testing.T) {
	retrieveAnimalFactsService := NewAnimalFactsService()
	retrieveAnimalFactsService.DogFactService = &MockDogFactService{}
	retrieveAnimalFactsService.CatFactService = &MockCatFactService{}

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
