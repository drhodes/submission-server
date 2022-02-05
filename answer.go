package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// from this submission server, send answers to the answer server
func sendAnswers(anonEdxId, labName, jsonAnswers string) error {
	answerServerUserId := os.Getenv("ANSWER_SERVER_USERID")
	answerServerPassword := os.Getenv("ANSWER_SERVER_PASSWD")
	answerServer := os.Getenv("ANSWER_SERVER")

	vals := url.Values{
		"edx-anon-id": {anonEdxId},
		"lab-answers": {jsonAnswers},
		"labname":     {labName},
	}
	log.Println(vals)

	endpoint := fmt.Sprintf("https://%s/submit-answers", answerServer)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(vals.Encode()))
	if err != nil {
		return Err(err, "Couldn't build request to send to answer server")
	}

	req.SetBasicAuth(answerServerUserId, answerServerPassword)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Err(err, "Had trouble sending request to the answer server")
	}

	if resp.StatusCode != 200 {
		log.Println(resp)
		return Err(nil, fmt.Sprintf("Got a bad response code: %d, %s", resp.StatusCode, resp))
	}
	return nil
}
