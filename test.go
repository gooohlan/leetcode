package main

func fibonacci() func() int {
	var f1, f2 int = 0, 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
}
