// /home/krylon/go/src/github.com/blicero/rolodex/common/logfacility/id.go
// -*- mode: go; coding: utf-8; -*-
// Created on 06. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-06 23:24:06 krylon>

// Package logfacility provides symbolic constants for the various parts
// of the application that might want to do logging.
package logfacility

//go:generate stringer -type=ID

// ID represents a part of the application that wants to write log messages.
type ID uint8

const (
	Common ID = iota
	LDAP
	GUI
)

func All() []ID {
	return []ID{
		Common,
		LDAP,
		GUI,
	}
} // func All() []ID
