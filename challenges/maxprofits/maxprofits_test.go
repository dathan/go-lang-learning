package challenge_test

import "testing"

/*
Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func TestMaxProcess(t *testing.T) {

	//	input := []int{7, 6, 4, 3, 1}

	type args struct {
		input []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"maxProfit", args{input: []int{1, 2, 3, 4, 5}}, 4},
		{"maxProfit", args{input: []int{7, 1, 5, 3, 6, 4}}, 7},
		{"maxProfit", args{input: []int{7, 6, 4, 3, 1}}, 0},
	}

	for _, maxProf := range tests {

		mp := maxProfit2(maxProf.args.input)
		t.Logf("Output: %d", mp)

		if mp != maxProf.want {
			t.Errorf("FAIL: %d is not what we want %d\n", mp, maxProf.want)
		}

	}

}

// remember the sort is the trade order.
func maxProfit2(prices []int) int {

	//O(n) Time complexity O(n) Space Complexity O(1)
	var maxprofit int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] { // if the previous price is less then the current price sum the prices
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}

func maxProfitWRONG(prices []int) int {
	// buy low sell high -- max your profit.
	// let's skip max numbers and mark the day.
	// a day 0 - 5
	maxProfit := 0
	var profitRate int = 0

	for i, buyAt := range prices {
		maxProfitRate := 0
		for j, sellAt := range prices {
			if i >= j { // buy date cannot be greater then sell date
				continue
			}

			diff := sellAt - buyAt
			diffDays := (j - i) + 1 // have to count the current day
			// want the most profit in the least amount of days.
			profitRate = diff / diffDays
			if profitRate > maxProfitRate {
				maxProfit += diff
				maxProfitRate = profitRate
			}
		}

	}

	return maxProfit

}
