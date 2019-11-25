/*
 * Copyright (c) 2019 The searKing authors. All Rights Reserved.
 *
 * Use of this source code is governed by a MIT-style license
 * that can be found in the LICENSE file in the root of the source
 * tree. An additional intellectual property rights grant can be found
 * in the file PATENTS.  All contributing project authors may
 * be found in the AUTHORS file in the root of the source tree.
 */

package main

import (
	"syscall"

	"github.com/searKing/go-issue-call-go-in-c-signal-handler/signal"
)

func main() {
	sig := syscall.SIGINT
	//signal.TestGoHandlerCalledByCSignal(sig)
	// Output:
	// fatal: morestack on g0
	// Got none signal in 1s, Time overseed.

	//signal.TestGoHandlerCalledByCSignalWithSigAltStack(sig, signal.SIGSTKSZ)
	// Output:
	// fatal: morestack on g0
	// Got none signal in 1s, Time overseed.


	signal.TestCHandlerCalledByCSignal(sig)
	// Output:
	// Got signal: interrupt
}
