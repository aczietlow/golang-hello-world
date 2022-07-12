package main

import "fmt"

type Bytes float64

const (
	KB Bytes = 1 << (10 * (iota + 1)) // 00000001 || 1 + 1
	MB                                // 00000010 || 2
	GB                                // 00000100 || 4
	TB
)

// Using custom Type and iota, write a function that accepts any number of bytes and displays it in a human readable format.
func main() {
	output := Bytes(100000024561234)
	fmt.Println(output.toString())
}

func (b Bytes) toString() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	}
	return fmt.Sprintf("%.2f B", b)
}
