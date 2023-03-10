package main

import (
	"encoding/json"
	"io"
	"net/http"

	"example.com/webservice/service"
)

func main() {
	mux := http.NewServeMux()
	animalFactsService := service.NewAnimalFactsService()

	mux.HandleFunc("/random-animal-facts", randomAnimalFactsHandler(animalFactsService))
	http.ListenAndServe(":8080", mux)
}

func randomAnimalFactsHandler(animalFactsService service.AnimalFactsService) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getRandomAnimalFacts(animalFactsService, w)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
	return handler
}

func getRandomAnimalFacts(animalFactsService service.AnimalFactsService, w http.ResponseWriter) {
	animalFacts, err := animalFactsService.RetrieveAnimalFacts()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writeSuccessfulResponse(w, animalFacts)
}

func writeSuccessfulResponse(w http.ResponseWriter, responseBody interface{}) {
	w.Header().Add("content-type", "application/json")
	encodeResponseAsJSON(responseBody, w)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
