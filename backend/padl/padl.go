// /home/krylon/go/src/github.com/blicero/rolodex/backend/ldap/ldap.go
// -*- mode: go; coding: utf-8; -*-
// Created on 07. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-08 21:56:36 krylon>

// Package padl uses an LDAP directory as its backend.
// I wanted to name the package ldap, which would have been the obvious name,
// but that clashes, of course, with go-ldap. So padl is ldap spelled backwards.
package padl

import (
	"log"

	"github.com/blicero/rolodex/common"
	"github.com/blicero/rolodex/common/logfacility"
	"github.com/go-ldap/ldap/v3"
)

// LDAPConnection encapsulates a connection to an LDAP directory server
type LDAPConnection struct {
	addr string      // nolint: unused
	log  *log.Logger // nolint: unused
	conn *ldap.Conn  // nolint: unused
}

func Connect(addr string) (*LDAPConnection, error) {
	var (
		err error
		l   = &LDAPConnection{addr: addr}
	)

	if l.log, err = common.GetLogger(logfacility.LDAP); err != nil {
		return nil, err
	} // else if l.conn, err = ldap.
} // func Connect(addr string) (*LDAPConnection, error)
