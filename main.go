// /home/krylon/go/src/github.com/blicero/rolodex/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 06. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-06 23:30:07 krylon>

package main

import (
	"fmt"

	"github.com/blicero/rolodex/common"
)

func main() {
	fmt.Printf("%s %s, built on %s\n",
		common.AppName,
		common.Version,
		common.BuildStamp.Format(common.TimestampFormatMinute))

}
