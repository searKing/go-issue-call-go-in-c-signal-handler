/*
 * Copyright (c) 2019 The searKing authors. All Rights Reserved.
 *
 * Use of this source code is governed by a MIT-style license
 * that can be found in the LICENSE file in the root of the source
 * tree. An additional intellectual property rights grant can be found
 * in the file PATENTS.  All contributing project authors may
 * be found in the AUTHORS file in the root of the source tree.
 */

package signal

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

/*
#include <signal.h>
#include "signal.cgo.h"

// Forward Declaration
extern void GoOnSignal(int signum, siginfo_t *info, void *context) ;
*/
import "C"
import (
	"os/signal"
)

const SIGSTKSZ int = int(C.SIGSTKSZ) // /* (128K)recommended stack size */
//
// === cgo hooks for user-provided Go callbacks, and enums ===
//

func GetGoOnSignal() unsafe.Pointer {
	return C.GoOnSignal
}

//export GoOnSignal
func GoOnSignal(signum C.int, info *C.siginfo_t, context unsafe.Pointer) {
}

// TestCHandlerCalledByCSignal calls c function in c signal handler
func TestCHandlerCalledByCSignal(sig syscall.Signal) {
	C.CGO_SetSig(C.int(sig), C.CGO_OnSignalHandler(C.CGO_OnSignal))
	ListenAndKill(sig)
	// Output:
	// Got signal: interrupt
}

// TestGoHandlerCalledByCSignal calls go function in c signal handler
func TestGoHandlerCalledByCSignal(sig syscall.Signal) {
	C.CGO_SetSig(C.int(sig), C.CGO_OnSignalHandler(GetGoOnSignal()))
	ListenAndKill(sig)
	// Output:
	// fatal: morestack on g0
	// Got none signal in 1s, Time overseed.
}

// TestGoHandlerCalledByCSignalWithSigAltStack calls go function in c signal handler, with sigaltstack called before sigaction
func TestGoHandlerCalledByCSignalWithSigAltStack(sig syscall.Signal, sigstksz int) {
	C.CGO_SigAltStack(C.int(sigstksz))
	C.CGO_SetSig(C.int(sig), C.CGO_OnSignalHandler(GetGoOnSignal()))
	ListenAndKill(sig)
	// Output:
	// fatal: morestack on g0
	// Got none signal in 1s, Time overseed.
}

func ListenAndKill(sig syscall.Signal) {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig)

	// simulate to send a SIGINT to this example test
	go func() {
		_ = syscall.Kill(syscall.Getpid(), sig)
	}()
	// Block until a signal is received.
	select {
	case s := <-c:
		fmt.Printf("Got signal: %s.\n", s)
	case <-time.After(time.Second):
		fmt.Printf("Got none signal in 1s, Time overseed.\n")
	}

}
