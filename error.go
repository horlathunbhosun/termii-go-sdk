package termiigo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIError includes the response from the Termii API and some HTTP request info
type APIError struct {
	//	Message        string        `json:"message,omitempty"`
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
	//URL            *url.URL      `json:"url,omitempty"`
	//Header         http.Header   `json:"header,omitempty"`
}

// APIError supports the error interface
func (apiErr *APIError) Error() string {
	ret, _ := json.Marshal(apiErr)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func newAPIError(resp *http.Response) *APIError {
	p, _ := ioutil.ReadAll(resp.Body)

	var termiiErrorResponse ErrorResponse
	_ = json.Unmarshal(p, &termiiErrorResponse)
	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		//Header:         resp.Header,
		Details: termiiErrorResponse,
		//URL:            resp.Request.URL,
	}
}
