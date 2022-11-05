package httpclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SetParams() url.Values {
	params := url.Values{}

	// For debug purposes
	// params.Add("dataInicial", "05/08/2021")

	log.Println(strings.NewReader(params.Encode()))
	return params
}

func PrepareRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)

	if err != nil {
		fmt.Printf("%s", err)
	}

	return req
}

func Request(req *http.Request) *http.Response {
	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return response
}
