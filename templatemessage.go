package termiigo

import "net/http"

type TemplateMessageService termiiServices

// TemplateData TemplateDataObject is a representation of a template data
type TemplateDataObject struct {
	ProductName string `json:"product_name"`
	Otp         int    `json:"otp"`
	ExpiryTime  string `json:"expiry_time"`
}

// TemplateRequest is a representation of a template request
type TemplateRequest struct {
	PhoneNumber string             `json:"phone_number"`
	DeviceID    string             `json:"device_id"`
	TemplateID  string             `json:"template_id"`
	APIKey      string             `json:"api_key"`
	Data        TemplateDataObject `json:"data"`
}

// SendTemplateMessage  func to send Template Message
func (templateMessageService *TemplateMessageService) SendTemplateMessage(request TemplateRequest) (map[string]interface{}, error) {

	// Set the endpoint
	endPoint := "send/template"

	// Set API key in the request
	request.APIKey = templateMessageService.client.apiKey

	// Send the HTTP request
	err, responseMap := templateMessageService.client.MakeRequest(http.MethodPost, request, endPoint)
	if err != nil {
		return nil, err
	}
	// Return the response and no error
	return responseMap, nil

}
