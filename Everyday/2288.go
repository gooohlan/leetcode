package Everyday

import (
    "fmt"
    "strconv"
    "strings"
)

func discountPrices(sentence string, discount int) string {
    strs := strings.Split(sentence, " ")
    d := 1 - float64(discount)/100
    for i, s := range strs {
        if len(s) > 1 && s[0] == '$' {
            price, err := strconv.Atoi(s[1:]) // 出错表示不是全数字
            if err == nil {
                strs[i] = fmt.Sprintf("$%.2f", float64(price)*d)
            }
        }
    }
    return strings.Join(strs, " ")
}
