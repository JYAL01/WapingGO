package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//direccion de envio del json
	url := "http://waping.es/api/send"

	//creamos el json

	var jsonData = []byte(`
	{
		"token":"1cea1c9b02c4447588a8a91e1935c950",
		"source": 50766233431,
		"destination": 50760634535 ,
		"type":"text",
		"body":{
			"text":"Hello Word!"
		}
	  }
	`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response header", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
