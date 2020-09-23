package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Game Start\n")
	firstWord := "り"
	start(firstWord)
}

func start(word string) {
	fmt.Printf("最初の文字は：%s\n", word)

	text := input()
	judge(text, word)
}

func input() string {
	fmt.Printf("文字を入力してください。\n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(text, word string) {
	textRune := []rune(text)
	firstWord := string(textRune[0])
	size := len(textRune)

	if firstWord != word {
		fmt.Printf("最初の文字が違います。")
		fmt.Printf("入力した最初の文字：%s\n", firstWord)
		fmt.Printf("もう一度入力してください。")
		start(word)
	}

	/*
	   	whiteList := []string{
	   		"あ", "い", "う", "え", "お",
	   		"か", "き", "く", "け", "こ",
	   		"さ", "し", "す", "せ", "そ",
	   		"た", "ち", "つ", "て", "と",
	   		"な", "に", "ぬ", "ね", "の",
	   		"は", "ひ", "ふ", "へ", "ほ",
	   		"ま", "み", "む", "め", "も",
	   		"や", "ゆ", "よ",
	   		"ら", "り", "る", "れ", "ろ",
	   		"わ", "を", "ん",

	   		"ア", "イ", "ウ", "エ", "オ",
	   		"カ", "キ", "ク", "ケ", "コ",
	   		"サ", "シ", "ス", "セ", "ソ",
	   		"タ", "チ", "ツ", "テ", "ト",
	   		"ナ", "ニ", "ヌ", "ネ", "ノ",
	   		"ハ", "ヒ", "フ", "ヘ", "ホ",
	   		"マ", "ミ", "ム", "メ", "モ",
	   		"ヤ", "ユ", "ヨ",
	   		"ラ", "リ", "ル", "レ", "ロ",
	   		"ワ", "ヲ", "ン",
	   	}
	   :/

	   	/*
	   		for _, value := range textRune {
	   			if !InArray(whiteList, value) {
	   				fmt.Printf("入力値が不正です。\n")
	   			}
	   		}
	*/

	lastWord := string(textRune[size-1])
	fmt.Printf("最後の文字：%s", lastWord)

	if lastWord == "ん" {
		fmt.Printf("最後の文字が「ん」です。\n")
		end()
	}

	start(lastWord)
}

/*
func InArray(whiteList []string, word string) {
	// value の 値を取得するためにfor _.を使用
	for _, value := range whiteList {
		if word != value {
			return false
		}
	}
	return true
}
*/

func end() {
	fmt.Printf("Game Over\n")
	os.Exit(0)
}
