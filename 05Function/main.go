package main

func add(a, b int) (res int) {
	res = a + b
	return
}
func main() {
	res := add(5, 2)
	print(res) // 7
}
