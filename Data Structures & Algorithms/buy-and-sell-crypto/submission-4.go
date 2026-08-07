func maxProfit(prices []int) int {
    l := 0
    profit := 0
    for r := 0; r < len(prices); r++{
        if prices[l] > prices[r]{
            l = r
        } else {
            profit = max(profit, prices[r] - prices[l])            
        }
    }
    return profit
}
