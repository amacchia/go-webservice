package catclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomCatFact_Success(t *testing.T) {
	catFactResponse := `{"data":["Some random fact"]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(catFactResponse))
	}))
	defer server.Close()
	channel := make(chan string)
	catFactService := NewCatClient()
	catFactService.catFactsApiUrl = server.URL

	go catFactService.GetRandomCatFact(channel)

	actualCatFact := <-channel
	if actualCatFact != "Some random fact" {
		t.Errorf("Expected cat fact to be Some random fact, but was %s", actualCatFact)
	}
}
