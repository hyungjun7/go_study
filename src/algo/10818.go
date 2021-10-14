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

	var min, max, tmp int = 0, 0, 0
	scanner.Scan()
	tmp, _ = strconv.Atoi(scanner.Text())
	min = tmp
	max = tmp
	for i := 1; i < size; i++ {
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())
		if tmp < min {
			min = tmp
		} else if tmp > max {
			max = tmp
		}
	}
	fmt.Print(min, max)
}
