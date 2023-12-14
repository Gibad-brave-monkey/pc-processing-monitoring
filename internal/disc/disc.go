package disc

import (
	"fmt"
	"syscall"
	"text/tabwriter"
)

func DiscProcess(w *tabwriter.Writer) {
	var stat syscall.Statfs_t
	syscall.Statfs("/", &stat)
	fmt.Fprintf(w, "Total Disk Space:\t%v GB\n", (stat.Blocks*uint64(stat.Bsize))/1024/1024/1024)
	fmt.Fprintf(w, "Free Disk Space:\t%v GB\n", (stat.Bfree*uint64(stat.Bsize))/1024/1024/1024)
}
