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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getRandomAnimalFactsHandler(animalFactsService service.AnimalFactsService) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		animalFacts := animalFactsService.RetrieveAnimalFacts()
		c.JSON(http.StatusOK, &animalFacts)
	}
	return handler
}

// TODO: Use table tests where applicable
// TODO: Add error scenarios
// TODO: Logging?
// TODO: Code coverage in CI build
// TODO: Remove dependency on gin library
// TODO: Debug program
