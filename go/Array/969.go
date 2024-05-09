package Array

import "fmt"

func pancakeSort(arr []int) []int {
    var res []int
    var sort func(int)
    
    sort = func(n int) {
        if n == 1 { // 执行完毕
            return
        }
        
        // 找出最大索引
        index := 0
        for i := 1; i < n; i++ {
            if arr[i] > arr[index] {
                index = i
            }
        }
        reverse969(arr, 0, index)
        res = append(res, index+1)
        reverse969(arr, 0, n-1)
        res = append(res, n)
        sort(n - 1)
    }
    
    sort(len(arr))
    return res
}

func reverse969(arr []int, i, j int) {
    for i < j {
        fmt.Println(i, j)
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}
