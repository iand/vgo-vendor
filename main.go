package main

import "github.com/iand/circuit"

func main() {
	b := circuit.Breaker{}
	println(b.Concurrency)
}

//go:generate stringer -type=Foo
type Foo int

const (
	FooOne Foo = iota
	FooTwo
	FooThree
)
