package main

import (
	"bufio"
	"fmt"
	"os"
)

func compare(resVal *int32, resKey *int32, value int32, convert int32) {
	if *resVal < value {
		*resVal = value
		*resKey = convert
	} else if *resVal == value {
		*resKey = 63
	}
}

//https://www.acmicpc.net/problem/1157
func main() {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str := line[:len(line)-1]
	var array = make([]int32, 26, 26)
	var upper, lower, value, key int32 = 65, 97, 0, 0
	for _, v := range str {
		if v < 91 {
			l := v - upper
			array[l]++
			compare(&value, &key, array[l], v)
		} else {
			u := v - lower
			array[u]++
			compare(&value, &key, array[u], v-32)
		}
	}
	fmt.Printf("%c", key)
}
