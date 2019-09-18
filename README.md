# rLog
remote syslog, can be used in embedded linux like arm cpu.  
How To Use

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

