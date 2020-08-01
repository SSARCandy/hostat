package main

import (
  "fmt"

  "github.com/fatih/color"
  "github.com/shirou/gopsutil/mem"
)

func main() {
  v, _ := mem.VirtualMemory()

  // almost every return value is a struct
  fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

  // convert to JSON. String() is also implemented
  fmt.Println(v)
  // Print with default helper functions
  color.Cyan("Prints text in cyan.")

  // A newline will be appended automatically
  color.Blue("Prints %s in blue.", "text")
}

