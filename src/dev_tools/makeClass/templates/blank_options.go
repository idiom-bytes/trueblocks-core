// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * This file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

package [{ROUTE}]Pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

type [{PROPER}]Options struct {
[{OPT_FIELDS}]
}

func (opts *[{PROPER}]Options) TestLog() {
[{TEST_LOGS}]	opts.Globals.TestLog()
}

func (opts *[{PROPER}]Options) ToCmdLine() string {
	options := ""
[{DASH_STR}][{POSITIONALS}]	options += fmt.Sprintf("%s", "") // silence go compiler for auto gen
	return options
}

func FromRequest(w http.ResponseWriter, r *http.Request) *[{PROPER}]Options {
	opts := &[{PROPER}]Options{}
[{DEFAULTS_API}]	for key, value := range r.URL.Query() {
		switch key {
[{REQUEST_OPTS}]		default:
			if !globals.IsGlobalOption(key) {
				opts.BadFlag = validate.Usage("Invalid key ({0}) in {1} route.", key, "[{ROUTE}]")
				return opts
			}
		}
	}
	opts.Globals = *globals.FromRequest(w, r)

	return opts
}

var Options [{PROPER}]Options

func [{PROPER}]FinishParse(args []string) *[{PROPER}]Options {
	// EXISTING_CODE
	// EXISTING_CODE
	return &Options
}
