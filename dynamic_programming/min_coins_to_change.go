package dynamic_programming

import "sort"

func MinCoinsToChange(amount int, coins []int) int {
	sort.Ints(coins) //sort the coins

	minCoins := -1 //initially, we don't have a solution

	for i := len(coins) - 1; i >= 0; i-- { //iterate from the biggest coin to the smallest
		if coins[i] > amount { //if the coin is bigger than the amount, skip it
			continue
		}
		if coins[i] == amount { //jackpot!
			return 1
		}

		totalOfCoins := 0
		amountLeft := amount //initially, the amount left is the amount itself

		for j := i; j >= 0; j-- {
			var nCoins = amountLeft / coins[j] //find the current number of coins
			amountLeft = amountLeft % coins[j] //find the amount left
			totalOfCoins += nCoins             //add the number of coins to the total
		}

		if amountLeft == 0 && minCoins != -1 && totalOfCoins < minCoins { //if there is no amount left, we found a solution
			totalOfCoins = minCoins
		}
	}
	return minCoins
}
