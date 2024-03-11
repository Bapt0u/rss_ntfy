package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func SendNotif(body string, url string) {
	srv := url

	// Data is formated to match Webhook json format
	data := []byte(fmt.Sprintf("{\"msg\": \"%s\"}", body))
	req, err := http.NewRequest(http.MethodPost, srv, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error while creating request: %e", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error while contacting server: %e", err)
	}

	resp.Body.Close()
}
