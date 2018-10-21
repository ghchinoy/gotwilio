package gotwilio

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// BaseURL is the Twilio Base URL for SMS
var BaseURL = "https://api.twilio.com/2010-04-01"

// Client is a Twilio client
type Client struct {
	AccountSID string
	AuthToken  string
	BaseURL    string
}

// NewClient creates a new Twilio client
func NewClient(sid, auth string) *Client {
	c := &Client{sid, auth, BaseURL}
	return c
}

// SendSMS sends an SMS with the given message
func (c *Client) SendSMS(from, to, message string) error {
	twiliourl := fmt.Sprintf("%s/Accounts/%s/Messages.json", c.BaseURL, c.AccountSID)
	//log.Println(twiliourl)
	client := &http.Client{}
	v := url.Values{}
	v.Set("Body", message)
	v.Set("From", from)
	v.Set("To", to)
	payload := v.Encode()
	//log.Printf("%+v", v)

	//payload := fmt.Sprintf("Body=%s\nFrom=%s\nTo=%s", message, from, to)
	req, err := http.NewRequest("POST", twiliourl, strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.AccountSID, c.AuthToken)
	req.Header.Set("accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(res.Status)
	log.Printf("%s", body)
	return nil
}
