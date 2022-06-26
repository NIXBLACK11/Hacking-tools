package main

import (
   "io/ioutil"
   "log"
   "fmt"
   "os"
   "net/http"
   "bufio"
)
func main() {
	file, err := os.Open("root_domains.txt")
    if err != nil {
        fmt.Println("open time error")
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        line :=scanner.Text()
		resp:=check(line)
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
    		log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		fmt.Println(sb)
	}
	func check(line string){
		resp, err := http.Get(line)
		if err != nil {
    	log.Fatalln(err)
		}
		return resp
	}
}