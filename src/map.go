package main

import "fmt"

// マップを作成

func main() {

	capitals := make(map[string]string)
	// インデックスを使用してマップに値を格納
	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンD.C."
	capitals["中国"] = "北京"

	// インデックスを使用して値を取り出す
	fmt.Println(capitals["日本"])
	fmt.Println(capitals["アメリカ"])
	fmt.Println(capitals["中国"])

	// 出力を見やすくするため空行を出力
	fmt.Println()

	// rangeを使っても要素が取り出せる
	for country, capital := range capitals {
		fmt.Println(country, capital)
	}

	_, ok := capitals["イギリス"]
	if ok {
		fmt.Println("登録済み")
	} else {
		fmt.Println("未登録")
	}

	_, ok = capitals["日本"]
	if ok {
		delete(capitals, "日本")
	} else {
		fmt.Println("日本は存在しないため削除できません")
	}

	// rangeを使っても要素が取り出せる
	for country, capital := range capitals {
		fmt.Println(country, capital)
	}
}
