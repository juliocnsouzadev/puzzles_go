package dynamic_programming

func MinCoinsToChange(amount int, coins []int) int {
	totalOfCoins := 0
	amountLeft := amount //initially, the amount left is the amount itself

	for i := len(coins) - 1; i >= 0; i-- { //iterate from the biggest coin to the smallest
		if coins[i] > amount { //if the coin is bigger than the amount, skip it
			continue
		}
		if coins[i] == amount { //jackpot!
			return 1
		}
		var nCoins = amountLeft / coins[i] //find the current number of coins
		amountLeft = amountLeft % coins[i] //find the amount left
		totalOfCoins += nCoins             //add the number of coins to the total

		if amountLeft == 0 { //if there is no amount left, we found a solution
			return totalOfCoins
		}
	}
	return -1 //if we got here, it means that we couldn't find a solution
}
