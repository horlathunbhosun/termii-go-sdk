package termiigo

import (
	"net/http"
)

type NumberApiService termiiServices

type NumberApiRequest struct {
	APIKey string `json:"api_key"`
	To     string `json:"to"`
	Sms    string `json:"sms"`
}

// NumberSend Api function to send message with Number Api
func (numberApiService *NumberApiService) NumberSend(request NumberApiRequest) (map[string]interface{}, error) {
	// Set the endpoint
	endPoint := "sms/number/send/"

	// Set API key in the request
	request.APIKey = numberApiService.client.apiKey

	// Send the HTTP request
	err, responseMap := numberApiService.client.MakeRequest(http.MethodPost, request, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil
}
