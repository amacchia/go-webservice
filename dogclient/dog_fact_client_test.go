package dogclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomDogFact_Success(t *testing.T) {
	dogFactResponse := `{"facts":["Some random fact"],"success":true}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/facts" {
			t.Errorf("Expected path to be /api/facts, but was %s", r.URL.String())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(dogFactResponse))
	}))
	defer server.Close()
	channel := make(chan string)
	dogFactClient := NewDogFactClient()
	dogFactClient.dogFactsApiServerUrl = server.URL

	go dogFactClient.GetRandomDogFact(channel)

	actualDogFact := <-channel
	if actualDogFact != "Some random fact" {
		t.Errorf("Expected dog fact to be Some random fact, but was %s", actualDogFact)
	}
}
