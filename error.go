package termiigo

import (
	"encoding/json"
)

// / APIResponse is a custom interface representing the necessary parts of an HTTP response
type APIResponse interface {
	GetStatusCode() int
	GetBody() []byte
}

// HTTPResponse implements the APIResponse interface
type HTTPResponse struct {
	StatusCode int
	Body       []byte
}

func (r *HTTPResponse) GetStatusCode() int {
	return r.StatusCode
}

func (r *HTTPResponse) GetBody() []byte {
	return r.Body
}

// APIError includes the response from the Termii API and some HTTP request info
type APIError struct {
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
}

// APIError supports the error interface
func (apiErr *APIError) Error() string {
	ret, _ := json.Marshal(apiErr)
	return string(ret)
}

// ErrorResponse represents an error response from the Paystack API server
type ErrorResponse struct {
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func newAPIError(resp APIResponse) *APIError {
	var termiiErrorResponse ErrorResponse

	if resp != nil {
		// Assuming you have a way to parse the response body into termiiErrorResponse
		// For example, if the response body is JSON, you can use json.Unmarshal(resp.GetBody(), &termiiErrorResponse)
		// Ensure you handle any errors during decoding
		// For simplicity, let's assume termiiErrorResponse is populated correctly
	}

	return &APIError{
		HTTPStatusCode: resp.GetStatusCode(),
		Details:        termiiErrorResponse,
	}
}
