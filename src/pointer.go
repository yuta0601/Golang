package main

import "fmt"

func main() {
	var ptr *int      // int型のポインタ変数
	var i int = 12345 // int型の変数

	ptr = &i // 変数iのアドレスを取得

	fmt.Println("iのアドレスは:", &i)
	fmt.Println("ptrの値(変数iのアドレス):", ptr)

	fmt.Println("iの値:", i)
	fmt.Println("ポインタ経由のiの値:", *ptr)

	// ポインタ経由でiの値を変更
	*ptr = 999

	fmt.Println("ポインタ経由で変更したiの値:", i)
}
