/*
 *  Copyright 2019 The searKing authors. All Rights Reserved.
 *
 *  Use of this source code is governed by a MIT-style license
 *  that can be found in the LICENSE file in the root of the source
 *  tree. An additional intellectual property rights grant can be found
 *  in the file PATENTS.  All contributing project authors may
 *  be found in the AUTHORS file in the root of the source tree.
 */
#ifndef CGO_SIGNAL_CGO_H_
#define CGO_SIGNAL_CGO_H_
#include <signal.h>
#ifdef __cplusplus
extern "C" {
#endif

// Callbacks Predefinations
typedef void (*CGO_OnSignalHandler)(int signum, siginfo_t *info, void *context);

int CGO_OnSignal(int signum, siginfo_t *info, void *context);
int CGO_SigAltStack(int sigStkSz);
int CGO_SetSig(int signum, CGO_OnSignalHandler onSignal);

#ifdef __cplusplus
}
#endif

#endif // CGO_SIGNAL_CGO_H_
