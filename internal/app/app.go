package appp

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Gibad-brave-monkey/pc-process-monitoring/internal/cpu"
	"github.com/Gibad-brave-monkey/pc-process-monitoring/internal/disc"
	"github.com/Gibad-brave-monkey/pc-process-monitoring/internal/mem"
)

func PrintSystemInfo(w *tabwriter.Writer) {
	fmt.Fprintf(w, "PID:\t%d\n", os.Getpid())
	// Информация о памяти
	mem.MemInfo(w)

	// Информация о CPU
	cpu.CpuProcess(w)

	// Информация о диске (для корня файловой системы)
	disc.DiscProcess(w)
}
