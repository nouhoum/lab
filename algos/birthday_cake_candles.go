package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	max := candles[0]
	for i := 1; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
		}
	}

	var count int32
	for i := 0; i < len(candles); i++ {
		if candles[i] == max {
			count += 1
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	candlesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	candlesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var candles []int32

	for i := 0; i < int(candlesCount); i++ {
		candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
		checkError(err)
		candlesItem := int32(candlesItemTemp)
		candles = append(candles, candlesItem)
	}

	result := birthdayCakeCandles(candles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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

/*
You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.

Example


The maximum height candles are  units high. There are  of them, so return .

Function Description

Complete the function birthdayCakeCandles in the editor below.

birthdayCakeCandles has the following parameter(s):

int candles[n]: the candle heights
Returns

int: the number of candles that are tallest
Input Format

The first line contains a single integer, , the size of .
The second line contains  space-separated integers, where each integer  describes the height of .

Constraints

Sample Input 0

4
3 2 1 3
Sample Output 0

2
Explanation 0

Candle heights are . The tallest candles are  units, and there are  of them.
*/
