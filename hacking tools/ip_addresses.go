package main

import (
   "io/ioutil"
   "fmt"
   "log"
   "net/http"
   "encoding/json"
   "os"
   "bufio"
)

func main() {
   resp, err_request := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
   if err_request != nil {
      log.Fatalln(err_request)
   }
      file, err := os.Create("output.txt")
    if err != nil {
           fmt.Println("File does not exists or cannot be created")
    }
	w := bufio.NewWriter(file)
    defer file.Close()
 	byteValue, _ := ioutil.ReadAll(resp.Body) 
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	// fmt.Printf("%T",result["prefixes"])
	for _, i := range result["prefixes"].([]interface{}) {
		m := i.(map[string]interface{})
		fmt.Fprintf(w, "%v\n", m["ip_prefix"])
	}
	w.Flush()
}