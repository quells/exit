# Exit Signals and Notifications

## Installation

```
go get github.com/quells/exit@latest
```

## exit.Signal{}

Safely close a channel from multiple goroutines to send a signal exactly once.
Designed for synchronizing graceful shutdown of multiple HTTP servers in the
same Go process.

## exit.Notification()

Returns a channel which yields OS-level interrupt or termination signals.

```go
select {
case sig := <-exit.Notification():
    log.Printf("received exit signal %v; will now exit", sig)
}

// <CTRL-C>
// TIMESTAMP: received exit signal interrupt; will now exit
```
