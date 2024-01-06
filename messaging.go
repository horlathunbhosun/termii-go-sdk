package termiigo

import (
	"errors"
	"net/http"
)

type MessagingService termiiServices
type Media struct {
	Url     string `json:"url"`
	Caption string `json:"caption"`
}
type SendMessageRequestToArray struct {
	APIKey  string   `json:"api_key"`
	To      []string `json:"to"`
	From    string   `json:"from"`
	Sms     string   `json:"sms"`
	Type    string   `json:"type"`
	Channel string   `json:"channel"`
	Media   Media    `json:"media"`
}

type SendMessageRequest struct {
	APIKey  string `json:"api_key"`
	To      string `json:"to"`
	From    string `json:"from"`
	Sms     string `json:"sms"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Media   Media  `json:"-"`
}

// SendMessage Api function to send message both when the to is string or multiple
func (messagingService *MessagingService) SendMessage(request interface{}) (map[string]interface{}, error) {
	// Set the endpoint
	endPoint := "sms/send"
	// Set API key in the request

	//check for the type to send before sending it
	switch req := request.(type) {
	case SendMessageRequest:

		req.APIKey = messagingService.client.apiKey

		// Send the HTTP request
		err, responseMap := messagingService.client.MakeRequest(http.MethodPost, req, endPoint)
		if err != nil {
			return nil, err
		}
		// Return the response and no error
		return responseMap, nil
	case SendMessageRequestToArray:
		req.APIKey = messagingService.client.apiKey

		if len(req.To) > 100 {
			return nil, errors.New("you can only send 100 numbers in the array for this method")
		}
		// Send the HTTP request
		err, responseMap := messagingService.client.MakeRequest(http.MethodPost, req, endPoint)
		if err != nil {
			return nil, err
		}
		// Return the response and no error
		return responseMap, nil

	default:
		return nil, errors.New("unsupported request type")
	}

}

// SendMessageBulk Api function to send message in Bulk
func (messagingService *MessagingService) SendMessageBulk(request SendMessageRequestToArray) (map[string]interface{}, error) {

	// Set the endpoint
	endPoint := "sms/send/bulk"

	// Set API key in the request
	request.APIKey = messagingService.client.apiKey

	//check the length if it more than 10000
	if len(request.To) > 10000 {
		return nil, errors.New("you can only send 10,000 numbers in the array for this sendBulk")
	}
	// Send the HTTP request
	err, responseMap := messagingService.client.MakeRequest(http.MethodPost, request, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil
}
