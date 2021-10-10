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

	var arr = make([]int, size)
	var countArr = make([]int, 10001)
	var resultArr = make([]int, len(arr))
	for i := 0; i < size; i++ {
		var tmp int
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())
		countArr[tmp]++
	}

	for i := 1; i < 10001; i++ {
		countArr[i] += countArr[i-1]
	}
	for _, v := range arr {
		fmt.Println(countArr[v] - 1)
		resultArr[countArr[v]-1] = v
		countArr[v] -= 1
	}
	fmt.Println(resultArr)

	//for i := 1; i < 10001; i++ {
	//	for j := 0; j < countArr[i]; j++ {
	//		_, _ = fmt.Fprintln(writer, i)
	//	}
	//}
}
