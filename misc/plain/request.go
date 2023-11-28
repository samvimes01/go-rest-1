package plain

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Start() {
	// define the timeout of the request
	var timeout time.Duration = 1000 * time.Millisecond

	// create the HTTP client
	client := &http.Client{
		Timeout: timeout,
	}

	// define the URL for the request
	var URL string = "https://jsonplaceholder.typicode.com/posts/1"

	// create the GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Panic(err)
	}

	// send the request to the server
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	body, _ := io.ReadAll(res.Body)
	// print out the response body
	fmt.Println("Response body: ", string(body))
	fmt.Println("======")
	// print out the response status code
	fmt.Println("Response Status Code: ", res.StatusCode)
	fmt.Println("======")
	// print out the response headers
	fmt.Println("Response Headers")
	for k, v := range res.Header {
		fmt.Println(k, " : ", v)
	}
}