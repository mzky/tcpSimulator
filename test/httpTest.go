package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"log"
)

var (
	contentTypeJson = "application/json; charset=utf-8"
	scheme          = "http://"
)

func main() {
	address := "127.0.0.1:9000/test"
	fmt.Println(Get(address))
}

func Get(url string) error {
	request := createRequest(url, "GET", "")
	return exec(request, url)
}

func Post(url string, body interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		log.Printf("POST %s error %s", url, err.Error())
		return err
	}
	request := createRequest(url, "POST", string(payload))
	return exec(request, url)
}

func exec(request *http.Request, url string) error {
	client := &http.Client{
		Timeout: 60 * 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("POST %s error %s", url, err.Error())
		return err
	}
	respStr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("respStr:%v \n", string(respStr))

	return err
}

func createRequest(url string, method string, body string) *http.Request {
	fmt.Println(scheme + url)
	request, e := http.NewRequest(method, scheme+url, strings.NewReader(body))
	if e != nil {
		log.Printf("create request error %s %s body is %s", url, method, body)
		return nil
	}
	request.Header.Set("Content-Type", contentTypeJson)

	return request
}
