package main

import (
	"fmt"
	"git.garena.com/shopee/loan-service/credit_backend/credit_framework/go_tls/g"
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
