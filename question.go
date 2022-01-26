package nowcoder

/*
假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

思路：一定是先买后买，那么从第一位开始计算买入卖出，遍历一次即可
步骤：
1、首位置为买入和卖出
2、遇到更小的点，则替换为买入
3、计算最大收益
*/

func GoodStock(prices []int) (buyIndex, sellIndex, profit int) {
    if len(prices) < 1 {
        return
    }
    var (
        curBuyIndex int
        curBuy      = prices[curBuyIndex]
    )
    for index, value := range prices {
        if value < curBuy { // 把最小值挪过去
            curBuy = value
            curBuyIndex = index
        }
        curProfit := value - prices[curBuyIndex]
        if curProfit > profit {
            profit = curProfit
            buyIndex = curBuyIndex
            sellIndex = index
        }
    }
    return
}
