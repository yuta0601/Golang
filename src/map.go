package main

import "fmt"

// マップを作成

var Capitals = make(map[string]string)

func main() {

	// Capitals := make(map[string]string)
	// インデックスを使用してマップに値を格納
	Capitals["日本"] = "東京"
	Capitals["アメリカ"] = "ワシントンD.C."
	Capitals["中国"] = "北京"

	// インデックスを使用して値を取り出す
	fmt.Println(Capitals["日本"])
	fmt.Println(Capitals["アメリカ"])
	fmt.Println(Capitals["中国"])

	// 出力を見やすくするため空行を出力
	fmt.Println()

	// rangeを使っても要素が取り出せる
	// for country, capital := range Capitals {
	// 	fmt.Println(country, Capitals)
	// }

	if isExist("イギリス") {
		fmt.Println("登録済み")
	} else {
		fmt.Println("未登録")
	}
	if isExist("日本") { // capitalsに日本が存在
		delete(Capitals, "日本")
	} else { // capitalsに日本が存在しない
		fmt.Println("日本は存在しないため削除できません")
	}

	// rangeを使っても要素が取り出せる
	for country, Capital := range Capitals {
		fmt.Println(country, Capital)
	}
}

func isExist(keyName string) bool {
	msg := false
	_, rst := Capitals[keyName]

	if rst {
		msg = true
	} else {
		msg = false
	}
	return msg
}
