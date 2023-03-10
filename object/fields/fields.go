// /home/krylon/go/src/github.com/blicero/rolodex/object/fields/fields.go
// -*- mode: go; coding: utf-8; -*-
// Created on 10. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-10 20:26:21 krylon>

// Package fields provides symbolic constants for enumerating the
// fields/attributes of a Person
package fields

//go:generate stringer -type=ID

// ID is the type to identify fields/attributes in a contact
type ID uint8

const (
	NameFirst ID = iota
	NameLast
	NameFull
	Street
	City
	PostalCode
	Country
	Phone
	Mobile
	Email
	Fax
)
