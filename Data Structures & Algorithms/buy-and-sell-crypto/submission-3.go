func maxProfit(prices []int) int {
    res := 0
    l := 0
    for r := 0; r < len(prices); r++{
        if prices[l] > prices[r]{
            // update l 
            l = r
        }
        res = max(res, prices[r]-prices[l])
    }
    return res
}
