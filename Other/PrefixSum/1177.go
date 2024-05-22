package PrefixSum

import "math/bits"

func canMakePaliQueries(s string, queries [][]int) []bool {
    sum := make([][26]int, len(s)+1)

    for i, c := range s {
        sum[i+1] = sum[i]
        sum[i+1][c-'a']++
    }

    ans := make([]bool, len(queries))

    for i, q := range queries {
        left, right, k, m := q[0], q[1], q[2], 0
        for j := 0; j < 26; j++ {
            m += (sum[right+1][j] - sum[left][j]) % 2
        }
        ans[i] = m/2 <= k
    }
    return ans
}

func canMakePaliQueries2(s string, queries [][]int) []bool {
    sum := make([][26]int, len(s)+1)

    for i, c := range s {
        sum[i+1] = sum[i]
        sum[i+1][c-'a']++
        sum[i+1][c-'a'] %= 2 // 0为偶数,1为奇数
    }

    ans := make([]bool, len(queries))

    for i, q := range queries {
        left, right, k, m := q[0], q[1], q[2], 0
        for j := 0; j < 26; j++ {
            if sum[right+1][j] != sum[left][j] {
                m++
            }
        }
        ans[i] = m/2 <= k
    }
    return ans
}

func canMakePaliQueries3(s string, queries [][]int) []bool {
    sum := make([][26]int, len(s)+1)

    for i, c := range s {
        sum[i+1] = sum[i]
        sum[i+1][c-'a'] ^= 1 // 0为偶数,1为奇数
    }

    ans := make([]bool, len(queries))

    for i, q := range queries {
        left, right, k, m := q[0], q[1], q[2], 0
        for j := 0; j < 26; j++ {
            m += sum[right+1][j] ^ sum[left][j]
        }
        ans[i] = m/2 <= k
    }
    return ans
}

func canMakePaliQueries4(s string, queries [][]int) []bool {
    sum := make([]int, len(s)+1)

    for i, c := range s {
        b := 1 << (c - 'a')
        sum[i+1] = sum[i] ^ b // 数变偶数，偶数变奇数
    }

    ans := make([]bool, len(queries))

    for i, q := range queries {
        left, right, k := q[0], q[1], q[2]
        m := bits.OnesCount(uint(sum[right+1] ^ sum[left]))
        ans[i] = m/2 <= k
    }
    return ans
}
