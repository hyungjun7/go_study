package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var size int

	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	var tmp int
	var arr = make([]string, size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())
		arr[i] = strconv.FormatInt(int64(tmp), 2)
		for j := len(arr[i]) - 1; j > -1; j-- {
			if string(arr[i][j]) == "1" {
				_, _ = fmt.Fprint(writer, (j-(len(arr[i])-1))*-1, " ")
			}
		}
	}
}
