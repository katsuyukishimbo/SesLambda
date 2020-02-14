package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// BouncedRecipient return struct
type BouncedRecipient struct {
	EmailAddress string `json:"emailAddress"`
}

// Bounce return struct
type Bounce struct {
	BounceType        string             `json:"bounceType"`
	BouncedRecipients []BouncedRecipient `json:"bouncedRecipients"`
}

// BounceMessage return struct
type BounceMessage struct {
	NotificationType string `json:"notificationType"`
	Bounce           Bounce `json:"bounce"`
}

const (
	// REQURL FIXME: modify logic to create API.
	REQURL = "http://localhost:3000/api/v1/bounce_emails/create"
)

func execute(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}

// Post return error
func Post(client *http.Client, values url.Values) error {
	req, err := http.NewRequest("POST", REQURL, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	// FIXME: setting access-token.
	req.Header.Set("Authorization", "Bearer 1234")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	execute(resp)

	return err
}

// FilterEmail return string
func FilterEmail(snsEvent events.SNSEvent) string {
	var emailAddress string = ""

	for _, record := range snsEvent.Records {
		snsMsg := record.SNS.Message
		bounceMsg := BounceMessage{}

		if err := json.Unmarshal([]byte(snsMsg), &bounceMsg); err != nil {
			log.Printf("could not unmarshal message: %v", snsMsg)
		}

		if bounceMsg.Bounce.BounceType == "Permanent" {
			for _, recepient := range bounceMsg.Bounce.BouncedRecipients {
				emailAddress = recepient.EmailAddress
			}
		} else if bounceMsg.Bounce.BounceType == "Transient" {
			log.Printf("not blacklisting transient bounce for emails: %v", bounceMsg.Bounce.BouncedRecipients)
		}
	}

	return emailAddress
}

func handler(snsEvent events.SNSEvent) error {
	values := url.Values{}
	values.Add("email", FilterEmail(snsEvent))

	client := &http.Client{}
	Post(client, values)

	return nil
}

func main() {
	lambda.Start(handler)
}
