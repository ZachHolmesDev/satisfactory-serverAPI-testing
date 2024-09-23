package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
  "function": "QueryServerState",
  "data": {}
}
		`
		// 	jsonData :=
		// 		`
		// {
		//   "function": "HealthCheck",
		//   "data": {
		//     "ClientCustomData": ""
		//   }
		// }
		// 		`

	// data
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

func main() {
	print("-- start --\n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	token := os.Getenv("SF_SERVER_TOKEN")
	fmt.Printf("env var is %v\n", token)
	// getInput()
	sendRequest()
	print("-- end --\n")
}
