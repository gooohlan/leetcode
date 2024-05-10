package Other

import "container/list"

type MyHashSet struct {
    data []list.List
}

const base = 997

func Constructor() MyHashSet {
    return MyHashSet{make([]list.List, base)}
}

func (this *MyHashSet) hash(key int) int {
    return key % base
}

func (this *MyHashSet) Add(key int) {
    if !this.Contains(key) {
        h := this.hash(key)
        this.data[h].PushBack(key)
    }
}

func (this *MyHashSet) Remove(key int) {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(int) == key {
            this.data[h].Remove(e)
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(int) == key {
            return true
        }
    }
    return false
}
