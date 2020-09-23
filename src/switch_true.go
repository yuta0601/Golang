package main

import "fmt"

func main() {
	// ループ：-2~2
	for i := -2; i <= 2; i++ {
		switch true {
		case i > 0:
			fmt.Println("+")
		case i < 0:
			fmt.Println("-")
		default:
			fmt.Println("その他")
		}
	}
}
