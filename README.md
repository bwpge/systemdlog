# systemdlog

A library for writing log priority levels to stdout for systemd's journal.

## Installation

```
go get -u github.com/bwpge/systemdlog@latest
```

## Usage

Example program:

```go
package main

import (
    sdlog "github.com/bwpge/systemdlog"
)

func main() {
    sdlog.Info("Hello, world!")
    sdlog.Warnf("This is a warning: %d is %s", 42, "the answer")
}
```

Output:

```
<6>Hello, world!
<4>This is a warning: 42 is the answer
```
