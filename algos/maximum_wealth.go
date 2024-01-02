package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for i := 0; i < len(accounts); i++ {
		currentAmount := 0
		for _, a := range accounts[i] {
			currentAmount += a
		}
		if currentAmount > maxWealth {
			maxWealth = currentAmount
		}
	}

	return maxWealth
}

func main() {
	var arr = [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(arr))
}
