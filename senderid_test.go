package termiigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

const apiKey = "TL2yxYKBgewEkKMNWlcIGVkkfAITCqLHc8VE4zTEcdpZGwGzC7lhwHn3I6AYjc"

//var client = NewClient(os.Getenv("TERMII_API_KEY"), nil)

func TestGetSenderId(t *testing.T) {
	// Mock HTTP server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate request parameters
		assert.Equal(t, "/sender-id", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, apiKey, r.URL.Query().Get("api_key"))

		// Respond with a mock JSON response
		mockResponse := map[string]interface{}{
			"current_page": 1,
			"data": []interface{}{
				map[string]interface{}{
					"sender_id":  "ACME Key",
					"status":     "unblock",
					"company":    "",
					"usecase":    "",
					"country":    "",
					"created_at": "2021-03-29 16:51:53",
				},
				map[string]interface{}{
					"sender_id":  "ACME Alert",
					"status":     "unblock",
					"company":    "",
					"usecase":    "",
					"country":    "",
					"created_at": "2021-03-29 16:51:09",
				},
			},
			"first_page_url": "https://termii.com/api/sender-id?page=1",
			"from":           1,
			"last_page":      47,
			"last_page_url":  "https://termii.com/api/sender-id?page=47",
			"next_page_url":  "https://termii.com/api/sender-id?page=2",
			"path":           "https://termii.com/api/sender-id",
			"per_page":       10,
			"prev_page_url":  "",
			"to":             15,
			"total":          704,
		}
		responseJSON, _ := json.Marshal(mockResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}))
	defer mockServer.Close()

	// Set TERMII_API_KEY to a mock API key for testing
	os.Setenv("TERMII_API_KEY", apiKey)
	defer os.Unsetenv("TERMII_API_KEY")

	// Create a client with the mock server's URL
	client := NewClient(apiKey, nil)
	client.SenderIDServiceR.client.baseUrl, _ = url.Parse(mockServer.URL)

	// Call the GetSenderId method
	senderIdResponse, err := client.SenderIDServiceR.GetSenderId()

	//fmt.Println(senderIdResponse)
	// Assert that there is no error
	assert.NoError(t, err)

	// Assert the expected response
	expectedResponse := map[string]interface{}{
		"current_page": 1,
		"data": []interface{}{
			map[string]interface{}{
				"sender_id":  "ACME Key",
				"status":     "unblock",
				"company":    "",
				"usecase":    "",
				"country":    "",
				"created_at": "2021-03-29 16:51:53",
			},
			map[string]interface{}{
				"sender_id":  "ACME Alert",
				"status":     "unblock",
				"company":    "",
				"usecase":    "",
				"country":    "",
				"created_at": "2021-03-29 16:51:09",
			},
		},
		"first_page_url": "https://termii.com/api/sender-id?page=1",
		"from":           1,
		"last_page":      47,
		"last_page_url":  "https://termii.com/api/sender-id?page=47",
		"next_page_url":  "https://termii.com/api/sender-id?page=2",
		"path":           "https://termii.com/api/sender-id",
		"per_page":       10,
		"prev_page_url":  "",
		"to":             15,
		"total":          704,
	}

	//responseJSON, _ := json.Marshal(senderIdResponse)

	assert.Equal(t, expectedResponse, senderIdResponse)
}
