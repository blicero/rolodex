// /home/krylon/go/src/github.com/blicero/rolodex/backend/padl/01_connect_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 08. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-10 15:30:47 krylon>

package padl

import (
	"os"
	"testing"
)

const urlEnv = "ROLODEX_LDAP_URL"

var (
	l *LDAPConnection
)

func TestConnect(t *testing.T) {
	var (
		addr string
		err  error
	)

	if addr = os.Getenv(urlEnv); addr == "" {
		t.Log("No URL was supplied for an LDAP directory, skipping tests.")
	} else if l, err = Connect(addr); err != nil {
		t.Errorf("Cannot connect to server %q: %s\n",
			addr,
			err.Error())
	}

	// l.conn.Debug = true
} // func TestConnect(t *testing.T)
