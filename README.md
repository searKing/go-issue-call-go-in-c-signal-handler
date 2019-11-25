# go-issue-call-go-in-c-signal-handler
This repository contains an example of the problem

# [test](https://github.com/searKing/go-issue-call-go-in-c-signal-handler/blob/master/signal/example_test.go)

+ ExampleTestCHandlerCalledByCSignal
```calls c function in c signal handler, passed```
+ ExampleTestGoHandlerCalledByCSignal
```calls go function in c signal handler, blocked forever```
+ ExampleTestGoHandlerCalledByCSignalWithSigAltStack
```calls go function in c signal handler, with sigaltstack called before sigaction, blocked forever```

# [binary](https://github.com/searKing/go-issue-call-go-in-c-signal-handler/blob/master/cmd/main.go)

Execute as a binary
+ ```signal.TestCHandlerCalledByCSignal``` is executed
+ ```signal.TestGoHandlerCalledByCSignal``` is commented out
+ ```signal.TestGoHandlerCalledByCSignalWithSigAltStack``` is commented out
 