package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/webservice/model"
	"example.com/webservice/service"
	"github.com/gin-gonic/gin"
)

type MockAnimalFactsService struct{}
type MockAnimalFactsServiceWithError struct{}

const expectedDogFact = "Dogs are cool."
const expectedCatFact = "Cats are cool."

func (mockAnimalFactsService *MockAnimalFactsService) RetrieveAnimalFacts() (*model.AnimalFacts, error) {
	return &model.AnimalFacts{DogFact: expectedDogFact, CatFact: expectedCatFact}, nil
}

func (mockAnimalFactsServiceWithError *MockAnimalFactsServiceWithError) RetrieveAnimalFacts() (*model.AnimalFacts, error) {
	return nil, errors.New("error retrieving animal facts")
}

func Test_getRandomAnimalFactsHandler(t *testing.T) {
	tests := []struct {
		name               string
		animalFactsService service.AnimalFactsService
		expectedStatusCode int
		expectedDogFact    string
		expectedCatFact    string
		wantErr            bool
	}{
		{
			name:               "Success",
			animalFactsService: &MockAnimalFactsService{},
			expectedStatusCode: http.StatusOK,
			expectedDogFact:    expectedDogFact,
			expectedCatFact:    expectedCatFact,
		},
		{
			name:               "Error returned from animal facts service",
			animalFactsService: &MockAnimalFactsServiceWithError{},
			expectedStatusCode: http.StatusInternalServerError,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setUpRouter()
			router.GET("/random-animal-facts", getRandomAnimalFactsHandler(tt.animalFactsService))
			httpRecorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/random-animal-facts", nil)

			router.ServeHTTP(httpRecorder, req)
			res, _ := io.ReadAll(httpRecorder.Body)
			animalFacts := model.AnimalFacts{}
			json.Unmarshal(res, &animalFacts)

			if tt.expectedStatusCode != httpRecorder.Code {
				t.Errorf("Expected %d but got %d", tt.expectedStatusCode, httpRecorder.Code)
			}
			if !tt.wantErr && tt.expectedDogFact != animalFacts.DogFact {
				t.Errorf("Expected %s but got %s", tt.expectedDogFact, animalFacts.DogFact)
			}
			if !tt.wantErr && tt.expectedCatFact != animalFacts.CatFact {
				t.Errorf("Expected %s but got %s", tt.expectedCatFact, animalFacts.CatFact)
			}
		})
	}
}

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}
