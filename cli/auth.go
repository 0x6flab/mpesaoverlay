// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/0x6flab/mpesaoverlay/pkg/mpesa"

func Token(sdk mpesa.SDK) error {
	token, err := sdk.Token()
	if err != nil {
		logError(err)

		return nil
	}

	logJSON(token)

	return nil
}
