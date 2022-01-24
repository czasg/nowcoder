package stringmath

func StringNumberAdd(x, y string) string {
    var (
        index   int                           // result 下标
        value   uint8                         // 对应为相加的结果
        further bool                          // 进位
        result  = make([]rune, len(x)+len(y)) // 结果值
    )
    xi, yi := len(x)-1, len(y)-1        // 获取个位值的下标
    for xi >= 0 || yi >= 0 || further { // 遍历
        value = 0
        if xi >= 0 && yi >= 0 { // 从 x、y 的个位值开始相加
            value = x[xi] + y[yi] - 96
        } else if xi >= 0 { // 只有 x 值了
            value = x[xi] - 48
        } else if yi >= 0 { // 只有 y 值了
            value = y[yi] - 48
        }
        if further { // 是否已经进位
            further = false
            value += 1
        }
        if value > 9 { // 是否需要进位
            further = true
            value -= 10
        }
        result[index] = rune(value + 48) // 赋值结果值
        index++
        xi--
        yi--
    }
    result = result[:index]              // 取结果值切片
    for i := 0; i < len(result)/2; i++ { // 倒序
        xxi := i
        yyi := len(result) - i - 1
        result[xxi], result[yyi] = result[yyi], result[xxi]
    }
    return string(result)
}
