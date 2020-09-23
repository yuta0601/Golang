package main

import (
	"fmt"
	"time"
)


func main() {
	// ここコメント
	t := time.Now()
	unixTime := t.Unix()

	fmt.Println(unixTime)
}
