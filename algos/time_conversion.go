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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	last := strings.TrimSpace(s)[len(s)-2:]
	ss := strings.Split(s, last)[0]
	splits := strings.Split(ss, ":")
	hours, _ := strconv.Atoi(ss[0:2])
	shours := ""
	if last == "PM" {
		if hours != 12 {
			hours += 12
		}
		shours = fmt.Sprint(hours)
	} else {
		if hours == 12 {
			shours = "00"
		} else {
			shours = splits[0]
		}
	}

	return fmt.Sprintf("%s:%s:%s", shours, splits[1], splits[2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
