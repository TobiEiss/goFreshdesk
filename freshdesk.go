package gofreshdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Freshdesk is representive of the client
type Freshdesk struct {
	context.Context
	Domain string
	API    string
}

// query the api
func (freshdesk *Freshdesk) query(typ interface{}, route string, body interface{}, header map[string]string) error {
	// create httpURL
	httpURL := fmt.Sprintf("https://%s.freshdesk.com%s", freshdesk.Domain, route)

	// create http-Context
	httpContext, cancelFunc := context.WithTimeout(freshdesk, 15*time.Second)
	defer cancelFunc()

	// build request
	request, err := func() (*http.Request, error) {
		if body != nil {
			bodyBytes, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}

			return http.NewRequest("POST", httpURL, bytes.NewBuffer(bodyBytes))
		}
		return http.NewRequest("GET", httpURL, nil)
	}()
	if err != nil {
		return err
	}

	// add api-key
	request.SetBasicAuth(freshdesk.API, "")
	request.Header.Add("Content-type", "application/json")

	// add header if necessary
	for key, value := range header {
		request.Header.Add(key, value)
	}

	// fire up request and unmarshal serverTime
	err = hTTPDo(httpContext, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}

		log.Println(response.StatusCode)

		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		if err := decoder.Decode(&typ); err != nil {
			return err
		}
		return nil
	})
	return err
}

// hTTPDo function runs the HTTP request and processes its response in a new goroutine.
func hTTPDo(ctx context.Context, request *http.Request, processResponse func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to processResponse.
	transport := &http.Transport{}
	client := &http.Client{Transport: transport}
	errorChannel := make(chan error, 1)

	// do request
	go func() { errorChannel <- processResponse(client.Do(request)) }()
	select {
	case <-ctx.Done():
		transport.CancelRequest(request)
		<-errorChannel // wait for processResponse function
		return ctx.Err()
	case err := <-errorChannel:
		return err
	}
}
