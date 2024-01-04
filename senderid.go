package termiigo

import (
	"fmt"
	"net/http"
)

type SenderIDService termiiServices

// SenderIdResponse holds the response from the Termii API for the sender id endpoint.
type SenderIdResponse struct {
	CurrentPage int `json:"current_page"`
	Data        []struct {
		SenderId  string `json:"sender_id"`
		Status    string `json:"status"`
		Company   string `json:"company"`
		UseCase   string `json:"usecase"`
		createdAt string `json:"created_at"`
	} `json:"data"`
	FirstPageURL string `json:"first_page_url"`
	From         int    `json:"from"`
	LastPage     int    `json:"last_page"`
	LastPageURL  string `json:"last_page_url"`
	NextPageURL  string `json:"next_page_url"`
	Path         string `json:"path"`
	PerPage      int    `json:"per_page"`
	PrevPageURL  string `json:"prev_page_url"`
	To           int    `json:"to"`
	Total        int    `json:"total"`
}

/*
// GetSenderId gets the list of sender ids
// https://developer.termii.com/sender-id
@params: apiKey
*/
func (senderIDr *SenderIDService) GetSenderId() (*SenderIdResponse, error) {
	senderIdResponse := &SenderIdResponse{}
	err := senderIDr.client.sendRequest(http.MethodGet, fmt.Sprintf("sender-id?api_key=%s", senderIDr.client.apiKey), nil, &senderIdResponse)
	if err != nil {
		return &SenderIdResponse{}, err
	}
	return senderIdResponse, nil
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

// SenderIdRequestResponse represents the structure of the response payload for sender ID requests.
type SenderIdRequestResponse struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

// RequestSenderId sends a request to the "sender-id/request" endpoint.
func (senderIDr *SenderIDService) RequestSenderId(senderIdRequest RequestSenderIdRequest) (*SenderIdRequestResponse, error) {
	// Initialize the response struct
	//var senderIdRequestResponse SenderIdRequestResponse
	senderIdRequestResponse := &SenderIdRequestResponse{}

	// Set the endpoint
	endPoint := "sender-id/request"

	// Set API key in the request
	senderIdRequest.APIKey = senderIDr.client.apiKey

	// Send the HTTP request
	err := senderIDr.client.sendRequest(http.MethodPost, endPoint, senderIdRequest, &senderIdRequestResponse)
	if err != nil {
		//fmt.Println("error from senderIdRequestResponse")
		//fmt.Println(err)
		return nil, err
	}

	// Return the response and no error
	return senderIdRequestResponse, nil
}
