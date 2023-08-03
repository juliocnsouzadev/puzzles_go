package dynamic_programming

func NumberOfWaysToMakeChange(wantedAmount int, coins []int) int {
	if wantedAmount <= 0 || coins == nil || len(coins) == 0 {
		return 1
	}
	var ways = make([]int, wantedAmount+1)
	ways[0] = 1 //for zero amount there is one way to make change
	for _, coin := range coins {
		for amount := 1; amount <= wantedAmount; amount++ {
			if coin <= amount {
				previousWay := ways[amount-coin]
				currentWay := ways[amount]
				ways[amount] = currentWay + previousWay
			}
		}
	}
	return ways[wantedAmount]
}
