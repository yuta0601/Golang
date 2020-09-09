package main

import "fmt"

// Calculator 演算インターフェイス型
type Calculator interface {
	// 関数の定義
	Calculate(a int, b int) int
}

// Add 足算型
type Add struct {
	// フィールドは持たない
}

// Calculate Add型にCaluculatorインターフェイスのCalculator関数を実装
func (x Add) Calculate(a int, b int) int {
	return a + b
}

// Sub はフィールドを持たない
type Sub struct {
	// フィールドは持たない
}

// Calculate Add型にCaluculatorインターフェイスのCalculator関数を実装
func (x Sub) Calculate(a int, b int) int {
	return a - b
}

func main() {
	var add Add
	var sub Sub

	// インターフェイス型を宣言
	var cal Calculator

	cal = add
	fmt.Println("和", cal.Calculate(1, 2))

	cal = sub
	fmt.Println("差", cal.Calculate(1, 2))
}
