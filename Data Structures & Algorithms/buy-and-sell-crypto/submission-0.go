func maxProfit(prices []int) int {
    l := 0
    res := 0
    for r := 0; r < len(prices); r++{
        if prices[l] >= prices[r]{
            l = r
        } else {
            res = max(res, prices[r] - prices[l])
        }
    }
    return res
}
