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

// ポインタを使わずPersonを受け取る
func (p Person) addHeight() {
	p.height += 10
}

// ポインタを使ってPersonを受け取る
func (p *Person) addWeight() {
	p.weight += 10
}

func main() {
	q := Person{height: 160, weight: 50}
	fmt.Println(q) // {160 50 }

	q.addHeight()
	q.addWeight()

	fmt.Println(q) // {160 60 }
}
