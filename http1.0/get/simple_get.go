package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func simpleGet() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get("http://localhost:8080" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
