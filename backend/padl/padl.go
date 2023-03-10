// /home/krylon/go/src/github.com/blicero/rolodex/backend/ldap/ldap.go
// -*- mode: go; coding: utf-8; -*-
// Created on 07. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-10 23:42:55 krylon>

// Package padl uses an LDAP directory as its backend.
// I wanted to name the package ldap, which would have been the obvious name,
// but that clashes, of course, with go-ldap. So padl is ldap spelled backwards.
package padl

import (
	"log"
	"time"

	"github.com/blicero/rolodex/backend/query"
	"github.com/blicero/rolodex/common"
	"github.com/blicero/rolodex/common/logfacility"
	"github.com/blicero/rolodex/object/fields"
	"github.com/go-ldap/ldap/v3"
)

var attributes = map[fields.ID]string{
	fields.NameFirst: "givenName",
	fields.NameLast:  "sn",
	fields.NameFull:  "cn",
	fields.Street:    "st",
	fields.City:      "l",
}

// LDAPConnection encapsulates a connection to an LDAP directory server
type LDAPConnection struct {
	addr string      // nolint: unused
	log  *log.Logger // nolint: unused
	conn *ldap.Conn  // nolint: unused
}

// Connect opens a connection to an LDAP directory.
func Connect(addr string) (*LDAPConnection, error) {
	var (
		err error
		l   = &LDAPConnection{addr: addr}
	)

	if l.log, err = common.GetLogger(logfacility.LDAP); err != nil {
		return nil, err
	} else if l.conn, err = ldap.DialURL(addr); err != nil {
		l.log.Printf("[ERROR] Failed to connect to directory server %s: %s\n",
			addr,
			err.Error())
		return nil, err
	}

	l.conn.SetTimeout(time.Second * 3)

	// I have not the slightest idea why, but NOT calling Start fixed a
	// weird problem where searching a directory not return all results
	// that should have been returned.
	// Given the name of the method and the documentation, I expected
	// it to be necessary to call Start before interacting with the server,
	// but not only does it work without call Start, it works BETTER.
	// Weird, but as long as it works, I'm not complaining.
	// l.conn.Start()

	return l, nil
} // func Connect(addr string) (*LDAPConnection, error)

// Close closes the connection. After calling Close, no more interactions with
// the server are possible.
func (c *LDAPConnection) Close() error {
	c.conn.Close()
	c.conn = nil
	return nil
} // func (c *LDAPConnection) Close() error

func (c *LDAPConnection) Search(q *query.Query) error {
	var (
		err  error
		qstr string
	)
} // func (c *LDAPConnection) Search(q *query.Query) error
