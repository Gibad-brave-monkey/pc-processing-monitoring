package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	appp "github.com/Gibad-brave-monkey/pc-process-monitoring/internal/app"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Заголовки столбцов
	fmt.Fprintf(w, "Metric\tValue\n")
	fmt.Fprintf(w, "------\t-----\n")

	// Бесконечный цикл с интервалом 5 секунда
	for {
		appp.PrintSystemInfo(w)
		w.Flush() // сброс буфера tabwriter
		time.Sleep(5 * time.Second)
		fmt.Println("------------------------")
	}
}
