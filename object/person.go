// /home/krylon/go/src/github.com/blicero/rolodex/object/person.go
// -*- mode: go; coding: utf-8; -*-
// Created on 06. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-09 16:28:12 krylon>

// Package object provides the data types to live in our address book.
package object

// Person is the central data type in the application.
// A Person has a name, phone number(s), email address(es), and so forth.
type Person struct {
	ID        int64
	NameFirst string
	NameLast  string
	Address   map[string]PostalAddress
	Phone     map[string]string
	Email     map[string]string
}

// FullName returns a Person's full name, in the form commonly used in
// Europe and America, "FirstName LastName".
func (p *Person) FullName() string {
	return p.NameFirst + " " + p.NameLast
} // func (p *Person) FullName() string

// PostalAddress describes a location one can send snail mail to.
// Usually this is a Person's home or place of business.
type PostalAddress struct {
	Street     string
	Number     string
	PostalCode string
	Country    string
	Additional []string
}
