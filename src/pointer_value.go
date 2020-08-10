package main

import "fmt"

func main() {
	// int型の変数を2つ宣言
	a, b := 1, 1

	// 関数に変数を渡す
	// a: 値渡し、b: 参照渡し
	double(a, &b)

	fmt.Println("値渡し", a)
	fmt.Println("参照渡し", b)
}

func double(x int, y *int) {
	x = x * 2
	*y = *y * 2
}
