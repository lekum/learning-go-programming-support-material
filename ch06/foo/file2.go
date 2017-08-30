package foo

var bar struct {
	x, y int
}

func quux() {
	bar = bar * bar
}
