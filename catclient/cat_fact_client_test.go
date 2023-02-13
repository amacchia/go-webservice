package catclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/webservice/model"
)

func Test_catFactClientImpl_GetRandomCatFact(t *testing.T) {
	tests := []struct {
		name                 string
		apiResponseCode      int
		apiResonseBody       string
		expectedCatFact      string
		expectedErrorMessage string
		shouldCauseError     bool
	}{
		{
			name:                 "Successfully retrieve cat fact",
			apiResponseCode:      http.StatusOK,
			apiResonseBody:       `{"data":["Some random fact"]}`,
			expectedCatFact:      "Some random fact",
			expectedErrorMessage: "",
			shouldCauseError:     false,
		},
		{
			name:                 "Non-200 resonse code from cat fact API",
			apiResponseCode:      http.StatusInternalServerError,
			apiResonseBody:       "{}",
			expectedCatFact:      "",
			expectedErrorMessage: "500 status code returned from cat fact API",
			shouldCauseError:     false,
		},
		{
			name:                 "Bad response body",
			apiResponseCode:      http.StatusOK,
			apiResonseBody:       "",
			expectedCatFact:      "",
			expectedErrorMessage: "error parsing cat fact response body",
			shouldCauseError:     false,
		},
		{
			name:                 "Error calling cat fact API",
			apiResponseCode:      http.StatusInternalServerError,
			apiResonseBody:       "",
			expectedCatFact:      "",
			expectedErrorMessage: "error retrieving cat fact",
			shouldCauseError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := serverSetup(tt.apiResponseCode, tt.apiResonseBody)
			defer server.Close()
			channel := make(chan model.AnimalFactResult)
			catFactClient := NewCatFactClient()
			catFactClient.catFactsApiUrl = server.URL
			if tt.shouldCauseError {
				catFactClient.catFactsApiUrl = ""
			}

			go catFactClient.GetRandomCatFact(channel)
			actualCatFactResult := <-channel

			if tt.expectedCatFact != actualCatFactResult.AnimalFact {
				t.Errorf("Expected cat fact to be %s, but was %s", tt.expectedCatFact, actualCatFactResult.AnimalFact)
			}
			if tt.expectedErrorMessage != "" && tt.expectedErrorMessage != actualCatFactResult.Error.Error() {
				t.Errorf("Expected error message to be %s, but was %s", tt.expectedErrorMessage, actualCatFactResult.Error.Error())
			}
		})
	}
}

func serverSetup(statusCode int, responseBody string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
	return server
}
