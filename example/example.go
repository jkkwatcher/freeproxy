package main

import (
	"fmt"

	"github.com/jkkwatcher/freeproxy"
)

func main() {
	gen := freeproxy.New()
	fmt.Println(gen.Get())
}
