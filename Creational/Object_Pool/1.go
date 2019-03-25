package main

import (
	"fmt"
	"go/types"
)

type Pool chan *types.Object

func New(total int) *Pool {
	p := make(Pool,total)
	for i := 0; i<total; i++ {
		p <- new(types.Object)
	}
	return &p
}

func main(){

	p := New(2)

	select {
	case obj := <- *p:
		fmt.Printf("%v",obj)
	default:
		return
	}

}
