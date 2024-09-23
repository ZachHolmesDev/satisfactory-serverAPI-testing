package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendRequest() {
	url := "https://192.168.0.207:7777/api/v1/"

	transport := &http.Transport{ // holds a pointer to configure http
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: transport}
	// reader := io.reader
	jsonData :=
		`
{
  "function": "HealthCheck",
  "data": {
    "ClientCustomData": ""
  }
}
		`
	// make the request
	resp, err := client.Post(
		url,
		"application/json;charset=utf-8",
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// read the respBody of response into respBody var
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respBodyJson bytes.Buffer
	err = json.Indent(&respBodyJson, respBody, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", respBodyJson.String())

}

// 	url := "https://192.168.0.207:7777/api/v1/"

// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 	}

// 	fmt.Printf("%v", response)

// }

// func sendRequest(fucntion string, dataInput string) {
// 	http.Get("")
// }

// func getInput() {
// 	var functionInput string
// 	var dataInput string

// 	fmt.Println("give the function:")
// 	fmt.Scanln(&functionInput)
// 	fmt.Println("give the data:")
// 	fmt.Scanln(&dataInput)
// }

func main() {
	print("-- start --\n")
	// getInput()
	sendRequest()
	print("-- end --\n")
}
