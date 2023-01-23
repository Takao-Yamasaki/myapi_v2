package main

import (
	"fmt"
)

// 構造体
type Person struct {
	height int
	weight int
	name   string
}

// メソッド
func (p *Person) Greeting() {
	fmt.Printf("Hello, my name is %s\n", p.name)
}

func main() {
	p := Person{name: "hsaki"}
	// p: レシーバー
	p.Greeting()
}
