package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'budgetShopping' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY bundleQuantities
 *  3. INTEGER_ARRAY bundleCosts
 */

func budgetShopping(n int32, bundleQuantities []int32, bundleCosts []int32) int32 {
	var qt, cost int32
	l := len(bundleCosts)

	budget := n
	for i := 0; i < l; i++ {
		//Get current bundle size
		currentBundleCost := bundleCosts[i]
		//Compute the quantity we can buy with the remaining budget
		quant := budget / currentBundleCost

		// If quantity is greater than zero we can afford this.
		// So update our budget and total quantity
		if quant > 0 {
			//Compute total cost
			totalCost := quant * currentBundleCost
			// Increase the cost
			cost += totalCost
			// Decrease the budget
			budget -= totalCost

			// Update the quantity
			numberOfBooks := bundleQuantities[i]
			qt += quant * numberOfBooks
		}
	}

	return qt
}

func computeNumberOfBooksRec(n int32, qt int32, cost int32, bundleQuantities []int32, bundleCosts []int32) int32 {
	// n = budget
	result := qt
	for i := 0; i < len(bundleCosts); i++ {
		// First compute the cost of adding an additional bundle
		currentCost := cost + bundleCosts[i]

		// Check if we are within the budget budget
		if currentCost <= n {
			//if so compute the new cost and quantity
			newQt := qt + bundleQuantities[i]
			newCost := cost + bundleCosts[i]
			// Continue: we try to an additional bundle
			newResult := computeNumberOfBooksRec(n, newQt, newCost, bundleQuantities, bundleCosts)
			if result < newResult {
				// We have more bundles. Update the result
				result = newResult
			}
		}
	}

	return result
}

func budgetShoppingRec(n int32, bundleQuantities []int32, bundleCosts []int32) int32 {
	// Write your code here
	return computeNumberOfBooksRec(n, 0, 0, bundleQuantities, bundleCosts)
}

func main() {
	fmt.Println("=======")
	var n int32 = 50
	bundleCosts := []int32{24, 20}
	bundleQuantities := []int32{20, 19}

	result1 := budgetShopping(n, bundleQuantities, bundleCosts)
	result2 := computeNumberOfBooksRec(n, 0, 0, bundleQuantities, bundleCosts)

	fmt.Println("Result 1=", result1)
	fmt.Println("Result 2=", result2)
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	bundleQuantitiesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var bundleQuantities []int32

// 	for i := 0; i < int(bundleQuantitiesCount); i++ {
// 		bundleQuantitiesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		bundleQuantitiesItem := int32(bundleQuantitiesItemTemp)
// 		bundleQuantities = append(bundleQuantities, bundleQuantitiesItem)
// 	}

// 	bundleCostsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

// 	var bundleCosts []int32

// 	for i := 0; i < int(bundleCostsCount); i++ {
// 		bundleCostsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		bundleCostsItem := int32(bundleCostsItemTemp)
// 		bundleCosts = append(bundleCosts, bundleCostsItem)
// 	}

// 	result := budgetShopping(n, bundleQuantities, bundleCosts)

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
