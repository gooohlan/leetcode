package WeeklyContest

import (
    "sort"
    "strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
    score := map[string]int{}
    for _, s := range positive_feedback {
        score[s] = 3
    }
    for _, s := range negative_feedback {
        score[s] = -1
    }
    
    type pair struct {
        id, score int
    }
    
    pairs := make([]pair, len(report))
    
    for i, r := range report {
        pairs[i] = pair{student_id[i], 0}
        
        for _, w := range strings.Split(r, " ") {
            pairs[i].score += score[w]
        }
    }
    
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].score > pairs[j].score || pairs[i].score == pairs[j].score && pairs[i].id < pairs[j].id
    })
    
    res := make([]int, k)
    
    for i, p := range pairs[:k] {
        res[i] = p.id
    }
    return res
}
