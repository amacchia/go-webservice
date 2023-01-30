package main

import (
	"net/http"

	"example.com/webservice/service"
	"github.com/gin-gonic/gin"
)

func main() {
	animalFactsService := service.NewAnimalFactsService()
	r := gin.Default()
	r.GET("/random-animal-facts", getRandomAnimalFactsHandler(animalFactsService))
	r.Run()
}

func getRandomAnimalFactsHandler(animalFactsService service.AnimalFactsService) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		animalFacts := animalFactsService.RetrieveAnimalFacts()
		c.JSON(http.StatusOK, &animalFacts)
	}
	return handler
}
