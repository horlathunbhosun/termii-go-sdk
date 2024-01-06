package termiigo

import (
	"fmt"
	"net/http"
)

type SenderIDService termiiServices

/*
// GetSenderId gets the list of sender ids
// https://developer.termii.com/sender-id
@params: apiKey
*/
func (senderIDr *SenderIDService) GetSenderId() (map[string]interface{}, error) {
	err, resp := senderIDr.client.MakeRequest(http.MethodGet, nil, fmt.Sprintf("sender-id?api_key=%s", senderIDr.client.apiKey))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
// Request Sender ID
// https://developer.termii.com/sender-id
*/
type RequestSenderIdRequest struct {
	APIKey   string `json:"api_key"`
	SenderID string `json:"sender_id"`
	Usecase  string `json:"usecase"`
	Company  string `json:"company"`
}

// RequestSenderId sends a request to the "sender-id/request" endpoint.
func (senderIDr *SenderIDService) RequestSenderId(senderIdRequest RequestSenderIdRequest) (map[string]interface{}, error) {

	// Set the endpoint
	endPoint := "sender-id/request"
	// Set API key in the request
	senderIdRequest.APIKey = senderIDr.client.apiKey
	// Send the HTTP request
	err, responseMap := senderIDr.client.MakeRequest(http.MethodPost, senderIdRequest, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil
}
