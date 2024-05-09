package main

import "container/list"

type entry struct {
	key, value int
}

const base = 997

type MyHashMap struct {
	data []list.List
}

func Constructors() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(entry); v.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(entry); v.key == key {
			return v.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(entry); v.key == key {
			this.data[h].Remove(e)
		}
	}
}
