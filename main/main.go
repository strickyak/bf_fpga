package main

import "fmt"
import "os"

import bf "github.com/strickyak/bf_fpga"

func compile(s string, o *bf.State) {
  i := 0
  for _, c := range s {
    var b byte
    switch c {
      case '<': b = 1
      case '>': b = 2
      case '+': b = 3
      case '-': b = 4
      case ',': b = 5
      case '.': b = 6
      default:
        continue
    }
    println(i, b)
    o.Prog[i] = b
    i++
  }
}


func main() {
  var o bf.State
  compile(os.Args[1], &o)
  for _, b := range o.Prog {
    fmt.Printf("%d ", b)
  }
  fmt.Printf("\n")
  for {
    ok := o.Step()
    if !ok { break }
    println("pc", o.PC, "tp", o.TP, "m", o.Tape[o.TP])
  }
}
