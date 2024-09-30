package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	// Print memory info
	fmt.Printf("Total: %v, Free: %v, UsedPercent: %f%%\n", v.Total, v.Free, v.UsedPercent)
}
