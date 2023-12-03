package main

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

/*
Given an integer array, return sum of array so that each element is unique by adding some numbers to duplicate elements so that sum of unique elements is minimum.

i.e., if all elements in the array are unique, return the sum. If some elements are duplicates, then increment them to make sure all elements are unique so that the sum of these unique elements is minimum.
*/

/*
 * Complete the 'getMinimumUniqueSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func getMinimumUniqueSum(arr []int32) int32 {
	// Write your code here
	// First sort the array. Make the problem easier
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	//take the first element
	prev := arr[0]
	//Since the array has been sorted we are sure the first element go the sum
	sum := prev
	for i := 1; i < len(arr); i++ {
		// Take the second element
		ei := arr[i]
		// if the previous (first) element is equal of greater than the current one
		// it is set the previous element plus 1
		if prev >= ei {
			ei = prev + 1
		}
		//the current is the next previous
		prev = ei
		//add the current ajusted element to the sum
		sum += ei
	}

	return sum
}

func main() {
	assertTrue(getMinimumUniqueSum([]int32{1, 2, 3}) == 6)
	assertTrue(getMinimumUniqueSum([]int32{2, 2, 3}) == 9)
	assertTrue(getMinimumUniqueSum([]int32{1, 1, 1, 2, 3, 5}) == 21)
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var arr []int32

// 	for i := 0; i < int(arrCount); i++ {
// 		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		arrItem := int32(arrItemTemp)
// 		arr = append(arr, arrItem)
// 	}

// 	result := getMinimumUniqueSum(arr)

// 	fmt.Fprintf(writer, "%d\n", result)

// 	writer.Flush()
// }

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

func assertTrue(expected bool) {
	if !expected {
		panic("bad response")
	}
}
