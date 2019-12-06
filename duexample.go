package main

import (
	"fmt"
	"github.com/ldav1s/go-disk-usage/du"
)

var KB = uint64(1024)

func main() {
	usage, _ := du.NewDiskUsage("C:\\")
	fmt.Println("Free:", usage.Free()/(KB*KB))
	fmt.Println("Available:", usage.Available()/(KB*KB))
	fmt.Println("Size:", usage.Size()/(KB*KB))
	fmt.Println("Used:", usage.Used()/(KB*KB))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
