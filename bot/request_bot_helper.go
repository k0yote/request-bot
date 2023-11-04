package bot

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type requestBotHelper struct {
	endpoints    []string
	customHeader *CustomHeader
}

func newRequestBotHelper(opts *RequestBotOptions) (*requestBotHelper, error) {
	var (
		endpoints = make([]string, 0)
		err       error
	)

	if opts != nil {
		endpoints, err = makeRequestList(opts.requestURL, opts.requestCount)
		if err != nil {
			return nil, err
		}
	}

	return &requestBotHelper{
		endpoints:    endpoints,
		customHeader: opts.CustomHeader,
	}, nil
}

func makeRequestList(requestURL string, requestCount int) ([]string, error) {
	endpoints := []string{}

	for i := 0; i < requestCount; i++ {
		endpoints = append(endpoints, requestURL)
	}

	return endpoints, nil
}

func botRequest(endpoints []string, customHeader *CustomHeader) ([]RequestCh, error) {
	done := make(chan RequestCh, len(endpoints))
	errch := make(chan error, len(endpoints))
	for _, endpoint := range endpoints {
		go func(endpoint string) {
			reqBody, _ := reqBody()
			b, err := request(endpoint, customHeader, reqBody)
			fmt.Println(string(b))
			if err != nil {
				errch <- err
				done <- RequestCh{}
				return
			}

			done <- RequestCh{
				Endpoint: endpoint,
				Data:     b,
			}
			errch <- nil
		}(endpoint)
	}

	botRequestArr := make([]RequestCh, 0)
	var errStr string
	for i := 0; i < len(endpoints); i++ {
		botRequestArr = append(botRequestArr, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}

	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}

	return botRequestArr, err
}

func request(endpoint string, customHeader *CustomHeader, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json")
	if customHeader != nil {
		req.Header.Add(customHeader.Key, customHeader.Value)
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
