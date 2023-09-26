// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Package cli provides the command line interface for the mpesaoverlay.
// The cli package is responsible for parsing the command line arguments and
// passing them to the appropriate function.
//
// The cli package is divided into subpackages:
//
//  1. auth: responsible for the GetToken command.
//
//  2. b2c: responsible for the B2CPayment command.
//
//  3. balance: responsible for the AccountBalance command.
//
//  4. c2b: responsible for the C2BRegisterURL and C2BSimulate commands.
//
//  5. express: responsible for the stkpush and stkpushquery commands.
//
//  6. qrcode: responsible for the generateqrcode command.
//
//  7. reversal: responsible for the Reversal command.
//
//  8. tax: responsible for the RemitTax command.
//
//  9. transaction: responsible for the TransactionStatus command.
//
//  10. log: responsible for the logError and logJSON functions.
package cli
