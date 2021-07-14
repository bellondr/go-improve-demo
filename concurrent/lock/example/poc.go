package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

const (
	First = "Fuck"
	Second = "What`s the"
)

func main() {
	s := "hello"
	p := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(p.Len)
	testCo()
}

func testCo(){
	var s string
	go func () {
		i := 1
		for {
			i = 1 - i
			if i == 0 {
				s = First
			} else {
				s = Second
			}
			time.Sleep(10)
		}
	}()

	for {
		fmt.Println(s)
		time.Sleep(10)
	}

}
