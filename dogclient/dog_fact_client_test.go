package dogclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/webservice/model"
)

func Test_dogFactClientImpl_GetRandomDogFact(t *testing.T) {
	tests := []struct {
		name                 string
		apiResponseCode      int
		apiResonseBody       string
		expectedDogFact      string
		expectedErrorMessage string
		shouldCauseError     bool
	}{
		{
			name:                 "Successfully retrieve dog fact",
			apiResponseCode:      http.StatusOK,
			apiResonseBody:       `{"facts":["Some random fact"],"success":true}`,
			expectedDogFact:      "Some random fact",
			expectedErrorMessage: "",
			shouldCauseError:     false,
		},
		{
			name:                 "Non-200 resonse code from dog fact API",
			apiResponseCode:      http.StatusInternalServerError,
			apiResonseBody:       "{}",
			expectedDogFact:      "",
			expectedErrorMessage: "500 status code returned from dog fact API",
			shouldCauseError:     false,
		},
		{
			name:                 "Bad response body",
			apiResponseCode:      http.StatusOK,
			apiResonseBody:       "",
			expectedDogFact:      "",
			expectedErrorMessage: "error parsing dog fact response body",
			shouldCauseError:     false,
		},
		{
			name:                 "Error calling dog fact API",
			apiResponseCode:      http.StatusInternalServerError,
			apiResonseBody:       "",
			expectedDogFact:      "",
			expectedErrorMessage: "error retrieving dog fact",
			shouldCauseError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := serverSetup(t, tt.apiResponseCode, tt.apiResonseBody)
			defer server.Close()
			channel := make(chan model.AnimalFactResult)
			dogFactClient := NewDogFactClient()
			dogFactClient.dogFactsApiServerUrl = server.URL
			if tt.shouldCauseError {
				dogFactClient.dogFactsApiServerUrl = ""
			}

			go dogFactClient.GetRandomDogFact(channel)
			actualDogFactResult := <-channel

			if tt.expectedDogFact != actualDogFactResult.AnimalFact {
				t.Errorf("Expected dog fact to be %s, but was %s", tt.expectedDogFact, actualDogFactResult.AnimalFact)
			}
			if tt.expectedErrorMessage != "" && tt.expectedErrorMessage != actualDogFactResult.Error.Error() {
				t.Errorf("Expected error message to be %s, but was %s", tt.expectedErrorMessage, actualDogFactResult.Error.Error())
			}
		})
	}
}

func serverSetup(t *testing.T, statusCode int, responseBody string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/facts" {
			t.Errorf("Expected path to be /api/facts, but was %s", r.URL.String())
		}

		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
	return server
}
