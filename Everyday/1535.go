package Everyday

func getWinner(arr []int, k int) int {
    mn := arr[0]
    win := 0
    for i := 1; i < len(arr) && win != k; i++ {
        if arr[i] > mn {
            mn = arr[i]
            win = 1
        } else {
            win++
        }
    }
    return mn
}
