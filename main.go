// +build go1.14

// Copyright 2020 The go-xn Authors. All rights reserved.
// Use of this source code is governed by a Zlib license
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/Pengxn/go-xn/src"
)

func main() {
	fmt.Println(app.Banner)

	app.Run()
}
