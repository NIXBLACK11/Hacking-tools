package main

import (
	"bytes"
	"fmt"

	"encoding/json"
	//"bufio"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
)

func main() {
	//Encode the data
	var name string
	fmt.Scanln(&name)
	postBody := []byte(`{
		"apiKey": "at_5h38PhPOm4SQ1KAzza7rTyXifU9Pf",
		"searchType": "current",
		"mode": "purchase",
		"punycode": true,
		"searchAfter": 1654349253,
		"basicSearchTerms": {
			"include": [
				"Pinterest, Inc."
			]
		}
	}`)
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://reverse-whois.whoisxmlapi.com/api/v2", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	// file2, err := os.Create("t2mama.txt")
	// if err != nil {
	// 	fmt.Println("File does not exists or cannot be created")
	// }
	// defer file2.Close()
	// w := bufio.NewWriter(file2)
	// fmt.Print(sb)
	// fmt.Fprintf(w, "%v\n", sb)
	// w.Flush()
	log.Printf(sb)
}
