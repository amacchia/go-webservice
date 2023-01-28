package model

type AnimalFacts struct {
	DogFact string `json:"dogFact"` // Use camelCase for response body
	CatFact string `json:"catFact"`
}
