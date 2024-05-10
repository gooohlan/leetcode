package DataStructure

import "container/heap"

type Twitter struct {
    users map[int]*user
    time  int
}

type user struct {
    id        int
    followed  map[int]*user
    tweetList *tweet
}

type tweet struct {
    id   int
    time int
    next *tweet
}

type tweetHeap []*tweet

func (h tweetHeap) Len() int { return len(h) }

func (h tweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }

func (h tweetHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *tweetHeap) Push(x interface{}) { *h = append(*h, x.(*tweet)) }

func (h *tweetHeap) Pop() (v interface{}) {
    old := *h
    *h, v = old[:len(old)-1], old[len(old)-1]
    return
}

func Constructor355() Twitter {
    return Twitter{
        users: make(map[int]*user),
    }
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    text := tweet{
        id:   tweetId,
        time: t.getTime(),
    }
    user := t.getUser(userId)
    text.next = user.tweetList
    user.tweetList = &text
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    user := t.getUser(userId)
    texts := make([]*tweet, 0, len(user.followed)+1)
    if user.tweetList != nil {
        texts = append(texts, user.tweetList)
    }
    
    for _, u := range user.followed {
        if u.tweetList != nil {
            texts = append(texts, u.tweetList)
        }
    }
    
    res := make([]int, 0, 10)
    textHeap := tweetHeap(texts)
    h := &textHeap
    heap.Init(h)
    
    for len(res) < 10 && h.Len() > 0 {
        text := heap.Pop(h).(*tweet)
        res = append(res, text.id)
        if text.next != nil {
            heap.Push(h, text.next)
        }
    }
    return res
}

func (t *Twitter) Follow(followerId int, followeeId int) {
    if followerId == followeeId {
        return
    }
    
    user1 := t.getUser(followerId)
    user2 := t.getUser(followeeId)
    user1.followed[followeeId] = user2
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
    if followerId == followeeId {
        return
    }
    user := t.getUser(followerId)
    _ = t.getUser(followeeId)
    delete(user.followed, followeeId)
}

func (t *Twitter) getTime() int {
    t.time++
    return t.time
}

func (t *Twitter) getUser(id int) *user {
    u, ok := t.users[id]
    if !ok {
        u = &user{
            id:       id,
            followed: make(map[int]*user),
        }
        t.users[id] = u
    }
    return u
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
