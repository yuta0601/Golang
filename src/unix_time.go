package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	unixTime := t.Unix()

	fmt.Println(unixTime)
}
