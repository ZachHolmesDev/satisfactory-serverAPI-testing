package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func sendRequest() {
	url := "https://192.168.0.207:7777/api/v1/"

	transport := &http.Transport{ // holds a pointer to configure http
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: transport}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// response, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("ERROR: ", err)
	// }

	fmt.Printf("%v", resp)

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
