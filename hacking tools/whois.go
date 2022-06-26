package main

import (
	"bufio"
	//"io"
	"fmt"
	"os/exec"
	"strings"

	//"log"
	"os"
	"time"
)

func main() {
	//Vimeo, Inc.
	fmt.Print("ENTER THE NUMBER OF PARAMETERS:")
	var n int
	fmt.Scanln(&	n)
	var name string
	fmt.Print("ENTER THE NUMBER OF STRING:")
	fmt.Scanln(&name)
	res := strings.Split(name, "|")
	file1, err := os.Open("root_domains.txt")
	if err != nil {
		fmt.Println("open me")
	}
	defer file1.Close()
	file2, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
	}
	defer file2.Close()
	w := bufio.NewWriter(file2)
	scanner := bufio.NewScanner(file1)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, ",")
		time.Sleep(1 * time.Second)
		out, e := exec.Command("whois", line).Output()
		if e != nil {
			fmt.Println()
		}
		o := string(out)
		for i := 0; i < n; i++ {
			if strings.Contains(o, res[i]) {
				fmt.Fprintf(w, "%v\n", line)
				fmt.Println(line)
				w.Flush()
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner me kuch")
	}
}