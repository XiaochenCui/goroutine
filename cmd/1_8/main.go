package main

import (
	"fmt"
	"github.com/huandu/goroutine"
)

func main() {
	id := goroutine.GoroutineId()
	fmt.Println(id)
}
