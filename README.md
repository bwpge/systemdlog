# systemdlog

A library for writing log priority levels to stdout for systemd's journal.

## Installation

```
go get github.com/bwpge/systemdlog@latest
```

## Usage

Example program:

```go
package main

import (
	log "github.com/bwpge/systemdlog"
)

func main() {
	log.Infof("hello, world")
}
```

Output:

```
<6>hello, world
```
