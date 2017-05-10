// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package version

import (
	"fmt"

	. "github.com/platinasystems/go"
	"github.com/platinasystems/go/goes"
	"github.com/platinasystems/go/goes/lang"
)

const (
	Name    = "version"
	Apropos = "print HEAD of source"
	Usage   = "version"
)

var Packages = func() []map[string]string { return []map[string]string{} }

type Interface interface {
	Apropos() lang.Alt
	Kind() goes.Kind
	Main(...string) error
	String() string
	Usage() string
}

func New() Interface { return cmd{} }

type cmd struct{}

func (cmd) Apropos() lang.Alt { return apropos }

func (cmd) Kind() goes.Kind { return goes.DontFork }

func (cmd) Main(args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("%v: unexpected", args)
	}
	for _, m := range append([]map[string]string{Package}, Packages()...) {
		fmt.Print(m["importpath"], ": ", m["version"], "\n")
	}
	return nil
}

func (cmd) String() string { return Name }
func (cmd) Usage() string  { return Usage }

var apropos = lang.Alt{
	lang.EnUS: Apropos,
}
