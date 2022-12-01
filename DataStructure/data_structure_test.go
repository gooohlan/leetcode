package DataStructure

import "testing"

func TestLRU(t *testing.T) {
	constructor := Constructor2(3)
	constructor.Put(1, 1)
	constructor.Put(2, 2)
	constructor.Put(2, 2222)
	constructor.Put(3, 3)
	constructor.Put(4, 4)
	constructor.Get(2)
	constructor.Put(5, 5)
}

func TestLFU(t *testing.T) {
	constructor := Constructor460(0)
	constructor.Put(0, 0)
	constructor.Put(2, 2)
	constructor.Put(1, 1111)
	constructor.Get(2)
	constructor.Put(3, 3)
}