package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("https://cbum-fitness.com/")
	if err != nil {
		err = errors.New("client.Get ERROR")
		os.Exit(1)
	}
	defer resp.Body.Close()

	req, err := http.NewRequest("GET", "https://cbum-fitness.com/", nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err = client.Do(req)

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	os.Exit(0)
}
