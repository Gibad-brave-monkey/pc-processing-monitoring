package cpu

import (
	"fmt"
	"syscall"
	"text/tabwriter"
)

func CpuProcess(w *tabwriter.Writer) {
	var rusage syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	fmt.Fprintf(w, "CPU Time:\t%v seconds\n", rusage.Utime.Sec+rusage.Stime.Sec)
}
