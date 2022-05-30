package main

func main() {
	numbers := []int{0,1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		if oddEven(n) {
			println(n, "is even")
		} else {
			println(n, "is odd")
		}
	}
}

func oddEven(n int) bool {
	return n%2 == 0
}

