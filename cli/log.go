// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
)

func logError(err error) {
	boldRed := color.New(color.FgRed, color.Bold)
	boldRed.Fprintf(os.Stderr, "\nerror: ")

	fmt.Fprintf(os.Stderr, "%s\n", color.RedString(err.Error()))
}

func logJSON(iList ...interface{}) {
	for _, i := range iList {
		m, err := json.Marshal(i)
		if err != nil {
			logError(err)

			return
		}

		pj, err := prettyjson.Format(m)
		if err != nil {
			logError(err)

			return
		}

		fmt.Fprintf(os.Stdout, "\n%s\n", string(pj))
	}
}
