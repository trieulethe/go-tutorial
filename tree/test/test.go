package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Test1 struct {
	value int
}

type Test2 struct {
	size int
	next *Test1
}

func main() {

	a := 2
	p := &a
	fmt.Println("p: ", p)
	fmt.Println("&p: ", &p)
	fmt.Println("*p: ", *p)
	*p = 3
	fmt.Println("a: ", a)
	fmt.Println()
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(*p))
	fmt.Println(reflect.TypeOf(&p))

	t1 := &Test1{}
	fmt.Println("t1: ", t1)
	p2 := unsafe.Pointer(t1)
	fmt.Println("pointer of t1:", p2)
	fmt.Println(reflect.TypeOf(p2))
	fmt.Println(reflect.TypeOf(&p2))
	*p2 = Test1{2}
	fmt.Println("t1: ", t1)
	fmt.Println("&t1: ", &t1)
	t2 := Test2{1, t1}
	fmt.Println("test 2: ", t2)

	t2.next = nil
	fmt.Println("test 2: ", t2)

}
