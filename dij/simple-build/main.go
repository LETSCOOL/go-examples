// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/letscool/lc-go/dij"
	"log"
)

type SampleApp struct {
	lib1 *SampleLib1 `di:"lib1"`
	lib2 *SampleLib2 `di:"lib2"`
}

type SampleLib1 struct {
	lib2 *SampleLib2 `di:"lib2"`
}

type SampleLib2 struct {
	val int `di:"val"`
}

func main() {
	ref := DependencyReference{"val": 123}
	app, err := BuildInstance(&SampleApp{}, &ref, "^")
	if err != nil {
		log.Fatal(err)
	}
	if app.lib2 != app.lib1.lib2 {
		log.Fatalf("incorrect injection, app.lib2(%v) != app.lib1.lib2(%v)\n", app.lib2, app.lib1.lib2)
	}
	if app.lib2.val != 123 {
		log.Fatalf("incorrect injection, app.lib2.val(%d) != 123\n", app.lib2.val)
	}
}
