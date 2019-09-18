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
	fmt.Println("teste do log intelbras")
	rlog.AppSyslog(rlog.Debug, "%s ======== App teste log Version %s\n", rlog.ThisFunction(), rlog.GetVersion())
	rlog.LoggerClose()

}
