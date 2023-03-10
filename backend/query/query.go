// /home/krylon/go/src/github.com/blicero/rolodex/backend/query/query.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 03. 2023 by Benjamin Walkenhorst
// (c) 2023 Benjamin Walkenhorst
// Time-stamp: <2023-03-10 20:53:04 krylon>

// Package query provides backend-agnostic queries to be used as an interface
// between the frontend and the various backends.
package query

import (
	"github.com/blicero/rolodex/object"
	"github.com/blicero/rolodex/object/fields"
)

// Query represents a query to a backend
type Query struct {
	Request map[fields.ID]string
	Result  []object.Person
}
