package main

import (
	"fmt"
	"github.com/go-tech-lab/framework_go_tls/g"
	"sync"
)

func main() {
	mp := sync.Map{}
	gp := g.G()
	mp.Store(gp, 12345)
	for i := 0; i <= 100; i++ {
		v, _ := mp.Load(gp)
		fmt.Printf("v = %v\n", v)
	}
}
