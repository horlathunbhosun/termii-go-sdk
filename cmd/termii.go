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
	//senderIdResponse, err := client.SenderIDServiceR.GetSenderId()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//dataConverted, err := json.Marshal(senderIdResponse)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//log.Printf("Response : %+v", string(dataConverted))
	//
	phoneBooksResponse, err := client.CampaignServiceR.GetPhoneBooks()
	if err != nil {
		fmt.Println(err)
		return
	}
	dataConvertedPhonebooks, err := json.Marshal(phoneBooksResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Response phonebook : %+v", string(dataConvertedPhonebooks))

	// RequestSenderId
	//requestBody := termiigo.RequestSenderIdRequest{
	//	SenderID: "Myadmission",
	//	Company:  "Myadmissionlink Educational Consult Limited",
	//	Usecase:  "Hi user, your myadmissionlink verification code is ",
	//}

	// senderIdRequestResponse, err := client.SenderIDServiceR.RequestSenderId(requestBody)

	// fmt.Println(senderIdRequestResponse)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// dataConvertedRequestSendID, err := json.Marshal(senderIdRequestResponse)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// log.Printf("Response : %+v", string(dataConvertedRequestSendID))

	//send Message with multiple numbers

	//payload := termiigo.SendMessageRequest{
	//	To:      "2349053573600",
	//	From:    "MEC", //senderId
	//	Sms:     "Hi Welcome to myadmissionlink test message",
	//	Type:    "plain",
	//	Channel: "generic",
	//}

	//payloadArrayNumber := termiigo.SendMessageRequestArray{
	//	To:      numbers,
	//	From:    "MEC", //senderId
	//	Sms:     "Hi Welcome to myadmissionlink test message",
	//	Type:    "plain",
	//	Channel: "generic",
	//}

	//messageResponse, err := client.MessagingServiceR.SendMessage(payload)
	//
	//fmt.Println(payload)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//dataConvertedRequestSendID, err := json.Marshal(messageResponse)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//log.Printf("Response : %+v", string(dataConvertedRequestSendID))

	payload := termiigo.PhoneBookRequest{
		PhonebookID:   "a3ec9a79-7c04-47b1-8d80-0dadeb14ca1b",
		PhonebookName: "MyadmissionLink",
	}
	//updatePhoneBook
	phoneBook, err := client.CampaignServiceR.UpdatePhoneBook(payload)

	fmt.Println(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataConvertedPhoneBook, err := json.Marshal(phoneBook)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Response phonebook: %+v", string(dataConvertedPhoneBook))

}
