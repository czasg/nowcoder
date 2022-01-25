package nowcoder

/* 大数相加
思路：根据 ascii 码进行常规加法计算
步骤：
1、确定最长字符串
2、从最长字符的个位开始，与另一个字符个位相加，然后计算进位
3、从个位循环到最高位，需要确保短字符不会越界
4、相加结束后需要确定最高位是否需要进位
*/
func AddStringNumber(x, y string) string {
    xi, yi := len(x), len(y)
    if xi > yi { // 确保 y 最大
        xi, yi = yi, xi
        x, y = y, x
    }

    var (
        xx, yy, pre uint8
        result      = make([]byte, yi+1)
    )
    for i := 0; i < yi; i++ {
        yy = y[yi-i-1] - 48 // 取 y 的个位值
        xx = 0
        if (xi - i - 1) >= 0 { // 确保 x 没有越界
            xx = x[xi-i-1] - 48
        }
        value := xx + yy + pre    // 当前值
        pre = value / 10          // 进位值
        value %= 10               // 进位后的值
        result[yi-i] = value + 48 // 结果值
    }

    if pre > 0 { // 进位则首位补 1
        result[0] = 49
        return string(result)
    }

    return string(result[1:])
}

/* 字符串翻转
思路：交换对应位置
步骤：
1、初始化字节数组
2、遍历交换
*/
func ReverseString(v string) string {
    result := []rune(v)
    for i := 0; i < len(v)/2; i++ { // 倒序
        m := i
        n := len(v) - i - 1
        result[m], result[n] = rune(v[n]), rune(v[m])
    }
    return string(result)
}
