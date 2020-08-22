package main

import "fmt"

func main() {
	// 配列を宣言
	values := [...]int{0, 1, 2, 3, 4}

	// スライスして関数を受け渡す
	double(values[:])

	fmt.Println(values)
}

func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}
