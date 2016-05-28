package bf

func Compile(s string, o *State) {
  var stack []byte
  var i byte
  for _, c := range s {
    var b byte
    switch c {
      case '<': b = 1
      case '>': b = 2
      case '+': b = 3
      case '-': b = 4
      case ',': b = 5
      case '.': b = 6
      case '[':
        stack = append(stack, i)
        b = 0  // Temporarily
      case ']':
        n := len(stack)
        j := stack[n-1]
        stack = stack[0:n-1]
        b = 0x80 | j
        o.Prog[j] = 0xC0 | (i+1)
      default:
        continue
    }
    println(i, b)
    o.Prog[i] = b
    i++
  }
}
