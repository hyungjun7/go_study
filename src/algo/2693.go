package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var size, maxLength = 0, 10
	_, _ = fmt.Scanf("%d", &size)

	var arr = make([][]int, size, size)

	for i := 0; i < size; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		arr[i] = make([]int, maxLength, maxLength)
		for j := 0; j < maxLength; j++ {
			arr[i][j], _ = strconv.Atoi(line[j])
		}
		sort.IntSlice(arr[i]).Sort()
		_, _ = fmt.Fprintln(writer, arr[i][7])
	}
}
