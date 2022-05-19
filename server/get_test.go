package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"testing"
)

func Test_GET_ReadAll(t *testing.T) {

	url := "http://localhost:8080/videos"

	req, _ := http.NewRequest("GET", url, nil)

    // pragmatic:reviews == cHJhZ21hdGljOnJldmlld3M
	//								(Base64)
	req.Header.Add("authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}