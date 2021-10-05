package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//백준
func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	temp := strings.Trim(line, " ")
	result := strings.Split(temp, " ")
	var cnt = 0
	for _, v := range result {
		if v == " " || v == "" || v == "\n" {
			continue
		} else {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)
}
