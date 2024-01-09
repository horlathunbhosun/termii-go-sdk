package termiigo

import (
	"fmt"
	"net/http"
)

type CampaignService termiiServices

type PhoneBookRequest struct {
	APIKey        string `json:"api_key"`
	PhonebookName string `json:"phonebook_name"`
	Description   string `json:"description,omitempty"`
	PhonebookID   string `json:"phonebook_id"`
}

// Fetch Phonebooks
// https://developer.termii.com/phonebook
// Request Type : GET
func (campaignService *CampaignService) GetPhoneBooks() (map[string]interface{}, error) {

	err, resp := campaignService.client.MakeRequest(http.MethodGet, nil, fmt.Sprintf("phonebooks?api_key=%s", campaignService.client.apiKey))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Create a Phonebook
// Endpoint : https://api.ng.termii.com/api/phonebooks
// Request Type : POST
func (campaignService *CampaignService) CreatePhoneBook(request PhoneBookRequest) (map[string]interface{}, error) {

	// Set the endpoint
	endPoint := "phonebooks"
	// Set API key in the request
	request.APIKey = campaignService.client.apiKey

	// Send the HTTP request
	err, responseMap := campaignService.client.MakeRequest(http.MethodPost, request, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil

}

//Update Phonebook
//Endpoint : https://api.ng.termii.com/api/phonebooks/{phonebook_id}
//Request Type : PATCH

func (campaignService *CampaignService) UpdatePhoneBook(request PhoneBookRequest) (map[string]interface{}, error) {
	// Set the endpoint
	endPoint := fmt.Sprintf("phonebooks/%s", request.PhonebookID)

	// Set API key in the request
	request.APIKey = campaignService.client.apiKey

	// Send the HTTP request
	err, responseMap := campaignService.client.MakeRequest(http.MethodPatch, request, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil

}
