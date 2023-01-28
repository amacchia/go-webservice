package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/webservice/model"
	"github.com/gin-gonic/gin"
)

type MockAnimalFactsService struct{}

func (mockAnimalFactsService *MockAnimalFactsService) RetrieveAnimalFacts() *model.AnimalFacts {
	return &model.AnimalFacts{DogFact: "Dogs are cool.", CatFact: "Cats are cool."}
}

func TestGetRandomAnimalFactsHandler(t *testing.T) {
	router := setUpRouter()
	router.GET("/random-animal-facts", getRandomAnimalFactsHandler(&MockAnimalFactsService{}))
	httpRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/random-animal-facts", nil)

	router.ServeHTTP(httpRecorder, req)
	res, _ := io.ReadAll(httpRecorder.Body)
	animalFacts := model.AnimalFacts{}
	json.Unmarshal(res, &animalFacts)

	if httpRecorder.Code != 200 {
		t.Errorf("Expected 200 but got %d", httpRecorder.Code)
	}
	if animalFacts.DogFact != "Dogs are cool." {
		t.Errorf("Expected Dogs are Cool. but got %s", animalFacts.DogFact)
	}
	if animalFacts.CatFact != "Cats are cool." {
		t.Errorf("Expected Cats are Cool. but got %s", animalFacts.CatFact)
	}
}

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}
