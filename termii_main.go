package termiigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	defaultHttpTimeout = 60 * time.Second
)

var baseUrl = os.Getenv("TERMII_BASE_URL")

// TermiiConfig holds Termii API configuration, including API key and base URL.
type termiiServices struct {
	client *Client
}

// Response represents arbitrary response data
type Response interface{}

// RequestValues aliased to url.Values as a workaround
type RequestValues url.Values

// MarshalJSON to handle custom JSON decoding for RequestValues
func (v RequestValues) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3)
	for k, val := range v {
		m[k] = val[0]
	}
	return json.Marshal(m)
}

type Client struct {
	common termiiServices // Reuse a single struct instead of allocating one for each service on the heap.
	client *http.Client

	// termii api key
	apiKey string

	// termii base url
	baseUrl *url.URL

	//Termii Services
	SenderIDServiceR        *SenderIDService
	MessagingServiceR       *MessagingService
	NumberApiServiceR       *NumberApiService
	TemplateMessageServiceR *TemplateMessageService
	CampaignServiceR        *CampaignService
}

func NewClient(apiKey string, httpClient *http.Client) *Client {

	//check if httpclient is nil
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHttpTimeout}
	}
	baseUrl, _ := url.Parse(baseUrl)
	fmt.Println(baseUrl.String())
	clientData := &Client{
		client:  http.DefaultClient,
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}
	fmt.Println(clientData)

	clientData.common.client = clientData
	clientData.SenderIDServiceR = (*SenderIDService)(&clientData.common)
	clientData.MessagingServiceR = (*MessagingService)(&clientData.common)
	clientData.NumberApiServiceR = (*NumberApiService)(&clientData.common)
	clientData.TemplateMessageServiceR = (*TemplateMessageService)(&clientData.common)
	clientData.CampaignServiceR = (*CampaignService)(&clientData.common)
	return clientData
}

func MapToJSON(mapData interface{}) []byte {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func (clientVal *Client) MakeRequest(requestMethod string, data interface{}, url string) (error error, response map[string]interface{}) {
	postData := MapToJSON(data)

	requestUrl, _ := clientVal.baseUrl.Parse(url)

	request, err := http.NewRequest(requestMethod, requestUrl.String(), bytes.NewBuffer(postData))
	if err != nil {
		return err, nil
	}

	// Set the request headers
	request.Header.Set("Content-Type", "application/json")

	// Perform the request
	responseReceived, err := clientVal.client.Do(request)
	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return err, nil
	}
	var result map[string]interface{}
	json.NewDecoder(responseReceived.Body).Decode(&result)
	return nil, result

}
