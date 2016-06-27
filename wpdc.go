// Copyright 2016 The wpdc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run mkdeprecatedlist.go

// Package wpdc (WordPress Deprecated Checker) checks depecreated functions
// and classes being used in a WordPress plugin.
package wpdc

import (
	"github.com/stephens2424/php/lexer"
	"github.com/stephens2424/php/token"
)

type Result struct {
	Line             int
	DeprecatedName   string
	DeprecatedSource string
}

func Check(content []byte) []Result {
	// Simplify scanning by checking identifiers only.
	lex := token.Subset(lexer.NewLexer(string(content)), token.Significant)

	var results []Result
	for item := lex.Next(); item.Typ != token.EOF; item = lex.Next() {
		if item.Typ != token.Identifier {
			continue
		}

		// Check previous item to make sure this is a function call or class
		// extending deprecated class.
		prev := lex.Previous()
		lex.Next()

		// It's okay to redeclare deprecated's identifier. Though it'd good
		// to make sure it's a method of there's a function_exists check.
		if prev.Typ == token.Function {
			continue
		}
		// TOOD(gedex): check if prev is `extends`.

		if _, ok := deprecated[item.Val]; ok {
			r := Result{
				Line:             item.Position().Line,
				DeprecatedName:   item.Val,
				DeprecatedSource: deprecated[item.Val],
			}
			results = append(results, r)
		}
	}

	return results
}
