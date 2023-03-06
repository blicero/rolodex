// /home/krylon/go/src/github.com/blicero/rolodex/backend/ldap/ldap.go
// -*- mode: go; coding: utf-8; -*-
// Created on 07. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-07 00:47:57 krylon>

// Package ldap uses an LDAP directory as its backend.
package ldap

import (
	"log"

	"github.com/go-ldap/ldap/v3"
)

// LDAPConnection encapsulates a connection to an LDAP directory server
type LDAPConnection struct {
	addr string      // nolint: unused
	log  *log.Logger // nolint: unused
	conn *ldap.Conn  // nolint: unused
}
