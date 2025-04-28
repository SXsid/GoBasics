package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1238901831)
	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
