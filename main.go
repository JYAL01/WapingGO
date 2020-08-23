package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func EnviarCampanas() {
	url := "http://waping.es/api/send"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`
	{
		"token":"1cea1c9b02c4447588a8a91e1935c950",
		"source":50766233431,
		"destination":50760634535,
		"type":"text",
		"body":{
			"text":"Hello Word"
		}
	  }
		
	`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
