// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package g exposes goroutine struct g to user space.
package g

import (
	"fmt"
	"runtime"
	"unsafe"
)

func getg() unsafe.Pointer

// G returns current g (the goroutine struct) to user space.
func G() unsafe.Pointer {
	// m1 macbook 必须必须执行一个任意操作，之后调用getg()值才TMD开始稳定下来，坑死个人了
	if runtime.GOOS == "darwin" {
		fmt.Print()
	}
	return getg()
}

func Gabc() unsafe.Pointer {
	return getg()
}
