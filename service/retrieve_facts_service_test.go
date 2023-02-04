package service

import (
	"errors"
	"testing"

	"example.com/webservice/model"
)

type MockDogFactService struct{}
type MockCatFactService struct{}
type MockDogFactServiceError struct{}
type MockCatFactServiceError struct{}

const expectedDogFact = "Some random dog fact"
const expectedCatFact = "Some random cat fact"
const expectedDogFactError = "dog fact error"
const expectedCatFactError = "cat fact error"

func (mockDogFactService *MockDogFactService) GetRandomDogFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{AnimalFact: expectedDogFact}
}

func (mockCatFactService *MockCatFactService) GetRandomCatFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{AnimalFact: expectedCatFact}
}

func (mockDogFactServiceError *MockDogFactServiceError) GetRandomDogFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{Error: errors.New(expectedDogFactError)}
}

func (mockCatFactServiceError *MockCatFactServiceError) GetRandomCatFact(factChannel chan<- model.AnimalFactResult) {
	factChannel <- model.AnimalFactResult{Error: errors.New(expectedCatFactError)}
}

func Test_animalFactsServiceImpl_RetrieveAnimalFacts(t *testing.T) {
	tests := []struct {
		name                   string
		animalFactsServiceImpl *animalFactsServiceImpl
		wantErr                bool
		expectedDogFact        string
		expectedCatFact        string
		expectedErrorMessage   string
	}{
		{
			name:                   "Success",
			animalFactsServiceImpl: buildAnimalFactsService(),
			wantErr:                false,
			expectedDogFact:        expectedDogFact,
			expectedCatFact:        expectedCatFact,
		},
		{
			name:                   "Dog fact client returns an error",
			animalFactsServiceImpl: buildAnimalFactsServiceWithDogFactClientError(),
			wantErr:                true,
			expectedErrorMessage:   expectedDogFactError,
		},
		{
			name:                   "Cat fact client returns an error",
			animalFactsServiceImpl: buildAnimalFactsServiceWithCatFactClientError(),
			wantErr:                true,
			expectedErrorMessage:   expectedCatFactError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualAnimalFacts, err := tt.animalFactsServiceImpl.RetrieveAnimalFacts()

			if tt.wantErr && err == nil {
				t.Error("Expected error but did not get one")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Expected nil error but an error %v was returned", err)
			}
			if tt.wantErr && (err.Error() != tt.expectedErrorMessage) {
				t.Errorf("Expected error with message %s, but was %s", err, "")
			}
			if !tt.wantErr && tt.expectedDogFact != actualAnimalFacts.DogFact {
				t.Errorf("Expected dog fact to be %s, but was %s", expectedDogFact, actualAnimalFacts.DogFact)
			}
			if !tt.wantErr && expectedCatFact != actualAnimalFacts.CatFact {
				t.Errorf("Expected cat fact to be %s, but was %s", expectedCatFact, actualAnimalFacts.CatFact)
			}
		})
	}
}

func buildAnimalFactsService() *animalFactsServiceImpl {
	retrieveAnimalFactsService := NewAnimalFactsService()
	retrieveAnimalFactsService.dogFactClient = &MockDogFactService{}
	retrieveAnimalFactsService.catFactClient = &MockCatFactService{}
	return retrieveAnimalFactsService
}

func buildAnimalFactsServiceWithDogFactClientError() *animalFactsServiceImpl {
	retrieveAnimalFactsService := buildAnimalFactsService()
	retrieveAnimalFactsService.dogFactClient = &MockDogFactServiceError{}
	return retrieveAnimalFactsService
}

func buildAnimalFactsServiceWithCatFactClientError() *animalFactsServiceImpl {
	retrieveAnimalFactsService := buildAnimalFactsService()
	retrieveAnimalFactsService.catFactClient = &MockCatFactServiceError{}
	return retrieveAnimalFactsService
}
