/*
 * Copyright (c) 2019 The searKing authors. All Rights Reserved.
 *
 * Use of this source code is governed by a MIT-style license
 * that can be found in the LICENSE file in the root of the source
 * tree. An additional intellectual property rights grant can be found
 * in the file PATENTS.  All contributing project authors may
 * be found in the AUTHORS file in the root of the source tree.
 */

package signal_test

import (
	"syscall"

	"github.com/searKing/go-issue-call-go-in-c-signal-handler/signal"
)

func ExampleTestCHandlerCalledByCSignal() {
	signal.TestCHandlerCalledByCSignal(syscall.SIGINT)
	// Output:
	// Got signal: interrupt.
}

func ExampleTestGoHandlerCalledByCSignal() {
	signal.TestGoHandlerCalledByCSignal(syscall.SIGINT)
	// Output:
	// Got signal: interrupt.
}

func ExampleTestGoHandlerCalledByCSignalWithSigAltStack() {
	signal.TestGoHandlerCalledByCSignalWithSigAltStack(syscall.SIGINT, signal.SIGSTKSZ)
	// Output:
	// Got signal: interrupt.
}