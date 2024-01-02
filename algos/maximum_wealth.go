package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	amount := 0
	for i := 0; i < len(accounts); i++ {
		currentAccounts := accounts[i]
		currentAmount := 0
		for j := 0; j < len(currentAccounts); j++ {
			currentAmount += currentAccounts[j]
		}
		if currentAmount > amount {
			amount = currentAmount
		}
	}

	return amount
}

func main() {
	var arr = [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(arr))
}
