package main

import (
	"encoding/json"
	"fmt"
	termiigo "github.com/horlathunbhosun/termii-go-package"
	"log"
	"os"
)

func main() {

	client := termiigo.NewClient(os.Getenv("TERMII_API_KEY"), nil)

	// GetSenderId
	senderIdResponse, err := client.SenderIDServiceR.GetSenderId()
	if err != nil {
		fmt.Println(err)
		return
	}
	dataConverted, err := json.Marshal(senderIdResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Response : %+v", string(dataConverted))

	// RequestSenderId
	requestBody := termiigo.RequestSenderIdRequest{
		SenderID: "Myadmission",
		Company:  "Myadmissionlink Educational Consult Limited",
		Usecase:  "Hi user, your myadmissionlink verification code is ",
	}

	senderIdRequestResponse, err := client.SenderIDServiceR.RequestSenderId(requestBody)

	fmt.Println(senderIdRequestResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataConvertedRequestSendID, err := json.Marshal(senderIdRequestResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Response : %+v", string(dataConvertedRequestSendID))

}
