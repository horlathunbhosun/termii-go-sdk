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
	SenderIDServiceR *SenderIDService
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
	return clientData
}

// sendRequest performs an HTTP request to the specified endpoint with the given parameters.

// APIError represents an error returned by the API
type APIErrorr struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *APIErrorr) Error() string {
	return e.Message
}

func MapToJSON(mapData interface{}) []byte {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func (clientVal *Client) MakePostRequest(requestMethod string, data interface{}, url string) (error error, response map[string]interface{}) {
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

//func (clientVal *Client) sendRequest(requestMethod, endpointUrl string, bodyParams interface{}, respData interface{}) error {
//	requestUrl, _ := clientVal.baseUrl.Parse(endpointUrl)
//
//	fmt.Println(requestUrl.String())
//	fmt.Println(bodyParams)
//
//	var reqBody io.Reader
//	if bodyParams != nil {
//		// Encode the request body to JSON
//		body, err := json.Marshal(bodyParams)
//		fmt.Println(bodyParams)
//
//		if err != nil {
//			return err
//		}
//		//change the body to bytes
//		reqBody = bytes.NewBuffer(body)
//	}
//
//	request, err := http.NewRequest(requestMethod, requestUrl.String(), reqBody)
//	if err != nil {
//		return err
//	}
//
//	// Set the request headers
//	request.Header.Set("Content-Type", "application/json")
//
//	// Perform the request
//	responseReceived, err := clientVal.client.Do(request)
//	if err != nil {
//		return err
//	}
//	//close the response body
//	defer responseReceived.Body.Close()
//
//	// Read the response body
//	return clientVal.decodeResponse(responseReceived, respData)
//}

//func checkStatusCode(statusCode int) error {
//	if statusCode != http.StatusOK && statusCode != http.StatusNoContent && statusCode != http.StatusCreated {
//		//return err
//		//\fmt.Errorf("error received, got status code %v", statusCode)
//	}
//	return nil
//}

//func mapstruct(data interface{}, v interface{}) error {
//	config := &mapstructure.DecoderConfig{
//		Result:           v,
//		TagName:          "json",
//		WeaklyTypedInput: true,
//	}
//	decoder, err := mapstructure.NewDecoder(config)
//	if err != nil {
//		return err
//	}
//	err = decoder.Decode(data)
//	return err
//}

//func (clientVal *Client) decodeResponse(httpResp *http.Response, value interface{}) error {
//	var resp Response
//	respBody, _ := ioutil.ReadAll(httpResp.Body)
//	err := json.Unmarshal(respBody, &resp)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(resp)
//	if status, _ := resp.(bool); !status || httpResp.StatusCode >= 400 {
//		return newAPIError(resp)
//	}
//
//	if data := resp; data != nil {
//		switch t := resp.(type) {
//		case map[string]interface{}:
//			return mapstruct(data, value)
//		default:
//			_ = t
//			return mapstruct(resp, value)
//		}
//	}
//	// if response data does not contain data key, map entire response to v
//	return mapstruct(resp, value)
//}
