package stringtool

// 大数相加 - 字符串数字相加
func StringNumberAdd(x, y string) string {
    xi, yi := len(x), len(y)
    if xi > yi { // 确保 y 最大
        xi, yi = yi, xi
        x, y = y, x
    }

    var (
        xx, yy, further uint8
        result          = make([]byte, yi+1)
    )
    for i := 0; i < yi; i++ {
        yy = y[yi-i-1] - 48 // 取 y 值
        xx = 0
        if (xi - i - 1) >= 0 { // 确保 x 没有越界
            xx = x[xi-i-1] - 48
        }
        value := xx + yy + further // 当前值
        further = value / 10       // 进位值
        value %= 10                // 进位后的值
        result[yi-i] = value + 48  // 结果值
    }

    if further > 0 { // 进位则首位补 1
        result[0] = 49
        return string(result)
    }

    return string(result[1:])
}
