package bibao

func add() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
