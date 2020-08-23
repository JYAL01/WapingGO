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
		"token":"dea0d11bff646db6797ee04ew2ecd96097f8b",
		"source": 999999999999,
		"destination": 00000000000 ,
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
