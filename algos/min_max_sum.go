package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	if len(arr) != 5 {
		panic("bad array length")
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var max int64 = 0
	var min int64 = 0

	for i := 0; i < len(arr)-1; i++ {
		if i+1 <= len(arr)-1 {
			max += int64(arr[i+1])
		}

		if i <= len(arr)-2 {
			min += int64(arr[i])
		}
	}

	fmt.Printf("%d %d", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// This second solution is less interessing
func miniMaxSum2(arr []int32) {
	// Write your code here
	if len(arr) != 5 {
		panic("bad array length")
	}
	sums := make([]int64, len(arr))

	for i := 0; i <= len(arr)-1; i++ {
		var sumi int64
		for j := 0; j <= len(arr)-1; j++ {
			if j != i {
				sumi += int64(arr[j])
			}
		}
		sums[i] = sumi
	}
	var min, max int64 = sums[0], sums[0]

	for i := 1; i <= len(arr)-1; i++ {
		c := int64(sums[i])
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	fmt.Printf("%d %d", min, max)
}
