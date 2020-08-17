// +build go1.14

// Copyright 2020 The Go-xn Authors. All rights reserved.
// Use of this source code is governed by a Zlib license
// that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/Pengxn/go-xn/src/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("Fail to Start Command app.", err.Error())
	}
}
