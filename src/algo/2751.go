package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	var reader = bufio.NewReader(os.Stdin)
	var size int

	_, _ = fmt.Fscanf(reader, "%d\n", &size)
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		_, _ = fmt.Fscanf(reader, "%d\n", &arr[i])
	}

	sort.IntSlice(arr).Sort()
	for i := 0; i < size; i++ {
		_, _ = fmt.Fprintf(writer, "%d\n", arr[i])
	}

	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
		}
	}(writer)
}
