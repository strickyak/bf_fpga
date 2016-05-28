package bf

const SZ = 64

type State struct {
  Prog  [SZ]byte
  Tape  [SZ]byte
  PC    byte
  TP    byte
}

// Returns false when done.
func (o *State) Step() bool {
  m := o.Tape[o.TP]
  m1 := m+1
  m9 := m-1
  m0 := (m == 0)

  tp := o.TP
  t1 := tp + 1
  t9 := tp - 1

  pc := o.PC
  p1 := pc + 1

  opcode := o.Prog[o.PC]
  branching := (opcode & 0xC0) != 0x00  // Branch or branch0.
  op_branch := (opcode & 0xC0) == 0xC0  // Branch always.
  op_branch0 := (opcode & 0xC0) == 0xC0  // Branch if m == 0.
  dest := opcode & 0x2F  // Low 6 bits, for branch destination.
  op := opcode & 0x07  // Low 3 bits, for non-branching op.

  op_left := !branching && op == 1
  op_right := !branching && op == 2
  op_incr := !branching && op == 3
  op_decr := !branching && op == 4
  op_input := !branching && op == 5
  op_output := !branching && op == 6

  next_pc := p1
  if op_branch || (op_branch0 && m0) {
    next_pc = dest
  }

  next_tp := tp
  if op_left { next_tp = t9 }
  if op_right { next_tp = t1 }

  enable_write := (op_incr || op_decr || op_input)
  next_m := m
  if op_incr { next_m = m1 }
  if op_decr { next_m = m9 }
  if op_input { next_m = o.Input() }

  if op_output { o.Output(m) }

  // Save them.
  o.PC, o.TP = next_pc, next_tp
  if enable_write {
    o.Tape[tp] = next_m
  }
  return (op != 0)
}

func (o *State) Input() byte {
  return 4  // Input is always 4.
}

func (o *State) Output(x byte) {
  println("output: ", int(x))
}
