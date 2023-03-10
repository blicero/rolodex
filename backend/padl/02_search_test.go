// /home/krylon/go/src/github.com/blicero/rolodex/backend/padl/02_search_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 08. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-10 10:42:39 krylon>

package padl

import (
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestSearch(t *testing.T) {
	if l == nil {
		t.SkipNow()
	}
	var (
		err error
		req *ldap.SearchRequest
		res *ldap.SearchResult
	)

	req = ldap.NewSearchRequest(
		"ou=contacts,dc=krylon,dc=net",
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		1000,
		false,
		"(sn=Simpson)",
		[]string{"givenName", "sn", "mail", "streetAddress", "l"},
		nil)

	if res, err = l.conn.SearchWithPaging(req, 4096); err != nil {
		t.Errorf("Failed to perform search on directory: %s\n",
			err.Error())
	} else if len(res.Entries) < 3 {
		t.Errorf("Unexpected number of results: %d (expected 3)",
			len(res.Entries))
	}

	// fmt.Printf("SearchResult: %s\n",
	// 	spew.Sdump(res))

	for _, entry := range res.Entries {
		// t.Logf("Entry %d: %v",
		// 	idx,
		// 	entry.GetAttributeValue("cn"))
		entry.PrettyPrint(2)
	}
} // func TestSearch(t *testing.T)
