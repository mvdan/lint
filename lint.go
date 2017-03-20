// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package lint

import (
	"go/token"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

// A Checker points out issues in a program.
type Checker interface {
	Check(*loader.Program, *ssa.Program) ([]Warning, error)
}

// Warning represents an issue somewhere in a source code file.
type Warning struct {
	Pos token.Pos
	Msg string
}
