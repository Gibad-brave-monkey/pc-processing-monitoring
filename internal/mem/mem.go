package mem

import (
	"fmt"
	"runtime"
	"text/tabwriter"
)

func MemInfo(w *tabwriter.Writer) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Fprintf(w, "Allocated Memory:\t%v MB\n", mem.Alloc/1024/1024)
}
