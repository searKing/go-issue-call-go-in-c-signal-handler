#include "signal.cgo.h"
#include <stdlib.h>
#include <string.h>
// record the existing Go signal handler.
static CGO_OnSignalHandler GoOnSignal;

int CGO_OnSignal(int signum, siginfo_t *info, void *context) {
  if (GoOnSignal) {
    GoOnSignal(signum, info, context);
  }
  return EXIT_SUCCESS;
}

int CGO_SigAltStack(int sigStkSz) {
  stack_t ss;
  sigaltstack(NULL, &ss);
  ss.ss_sp = malloc(sigStkSz); // SIGSTKSZ
  ss.ss_size = sigStkSz;
  ss.ss_flags = 0;
  return sigaltstack(&ss, NULL);
}

int CGO_SetSig(int signum, CGO_OnSignalHandler onSignal) {
  struct sigaction sa;
  memset(&sa, 0, sizeof(sa));
  sigaction(signum, NULL, &sa);
  GoOnSignal = sa.sa_sigaction;

  sa.sa_flags = sa.sa_flags & (~SA_SIGINFO);
  sa.sa_flags = sa.sa_flags | SA_ONSTACK | SA_RESTART | SA_SIGINFO;
  sa.sa_sigaction = onSignal;
  return sigaction(signum, &sa, NULL);
}
