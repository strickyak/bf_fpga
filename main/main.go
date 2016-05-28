// go run main/main.go  '++++[>+++[>+<-]<-]>>.'

package main

import "fmt"
import "os"

import bf "github.com/strickyak/bf_fpga"

func main() {
  var o bf.State
  bf.Compile(os.Args[1], &o)
  for i, b := range o.Prog {
    fmt.Printf("[%02x]%02x ", i, b)
  }
  fmt.Printf("\n")
  for {
    ok := o.Step()
    if !ok { break }
    println("pc", o.PC, "tp", o.TP, "m", o.Tape[o.TP])
  }
}
