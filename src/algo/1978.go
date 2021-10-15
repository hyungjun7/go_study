package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var size = 0
	_, _ = fmt.Scanf("%d", &size)
	scanner.Scan()
	var arr = strings.Split(scanner.Text(), " ")
	var result = 0
	for i := 0; i < size; i++ {
		tmp, _ := strconv.Atoi(arr[i])
		var count = 0
		for j := 2; j <= tmp; j++ {
			if tmp%j == 0 {
				count++
			}
		}
		if count == 1 {
			result++
		}
	}
	fmt.Println(result)
}
