package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://redmine.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(response.Status))
}
