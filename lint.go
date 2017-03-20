// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package lint

import (
	"go/token"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

type Checker interface {
	Check(*loader.Program, *ssa.Program) ([]Warning, error)
}

type Warning struct {
	Pos token.Pos
	Msg string
}
