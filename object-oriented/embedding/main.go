package main

import "fmt"

type Bird struct {
	featherLength  int
	classification string
}

type Pigeon struct {
	Bird
	featherLength float64
	Name          string
}

func main() {
	p := Pigeon{Name: "TweeTy"}
	p.featherLength = 3.14
	fmt.Println(p)
}
