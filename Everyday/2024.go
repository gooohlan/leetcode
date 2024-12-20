package Everyday

func maxConsecutiveAnswers(answerKey string, k int) int {
    cnt := [2]int{}
    left := 0
    var ans int
    for right, c := range answerKey {
        cnt[c>>1&1]++
        for cnt[0] > k && cnt[1] > k {
            cnt[answerKey[left]>>1&1]--
            left++
        }
        ans = max(ans, right-left+1)
    }
    return ans
}
