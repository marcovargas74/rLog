# rLog
Remote syslog, can be used in embedded linux like arm cpu.  

## Install

- Make sure you have [Go](https://golang.org/doc/install) installed and have set your [GOPATH](https://golang.org/doc/code.html#GOPATH).
- Install rLog:

```sh
go get -u github.com/marcovargas74/rLog
```

##Example

```
package main

import (
	"fmt"

	rlog "github.com/marcovargas74/rLog"
)

func init() {
	rlog.Clear()
	rlog.StartLogger(true, rlog.Debug|rlog.Local0, "172.31.11.162:514")
	rlog.SetPrintLocal(true)
}

func main() {
	fmt.Println("test of log ")
	rlog.AppSyslog(rlog.Debug, "%s ======== App test log Version %s\n", rlog.ThisFunction(), rlog.GetVersion())
	rlog.LoggerClose()

}
```
