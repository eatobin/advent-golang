// Struct from day05:

type IntCode struct {
	input   int
	output  int
	pointer int
	memory  [678]int
}

// opcode from day05:

func opcode(icP *IntCode, instruction *[5]int) int {
  instruction = pad5(icP.memory[icP.pointer], instruction)
  switch instruction[4] {
  case 1:
    icP.memory[aParam(icP, instruction)] = bParam(icP, instruction) + cParam(icP, instruction)
    icP.pointer += 4
    return 1
  case 2:
    icP.memory[aParam(icP, instruction)] = bParam(icP, instruction) * cParam(icP, instruction)
    icP.pointer += 4
    return 1
  case 3:
    icP.memory[cParam(icP, instruction)] = icP.input
    icP.pointer += 2
    return 1
  case 4:
    icP.output = cParam(icP, instruction)
    icP.pointer += 2
    return 1
  case 5:
    if cParam(icP, instruction) != 0 {
      icP.pointer = bParam(icP, instruction)
    } else {
      icP.pointer += 3
    }
    return 1
  case 6:
    if cParam(icP, instruction) == 0 {
      icP.pointer = bParam(icP, instruction)
    } else {
      icP.pointer += 3
    }
    return 1
  case 7:
    if cParam(icP, instruction) < bParam(icP, instruction) {
      icP.memory[aParam(icP, instruction)] = 1
    } else {
      icP.memory[aParam(icP, instruction)] = 0
    }
    icP.pointer += 4
    return 1
  case 8:
    if cParam(icP, instruction) == bParam(icP, instruction) {
      icP.memory[aParam(icP, instruction)] = 1
    } else {
      icP.memory[aParam(icP, instruction)] = 0
    }
    icP.pointer += 4
    return 1
  default:
    return 0
  }
}

// Struct from intcode:

type IntCode struct {
	Input        int
	Output       []int
	Phase        int
	Pointer      int
	RelativeBase int // added in day09
	Memory       map[int]int
	IsStopped    bool
	DoesRecur    bool
}

// opcode from intcode:

func (icP *IntCode) OpCode() int {
	if icP.IsStopped {
		return 0
	} else {
		instruction := pad5(icP.Memory[icP.Pointer])
		if instruction['d'] == 9 {
			icP.IsStopped = true
			return 0
		} else {
			switch instruction['e'] {
			case 1:
				icP.Memory[icP.aParam(instruction)] = icP.bParam(instruction) + icP.cParam(instruction)
				icP.Pointer += 4
				return 1
			case 2:
				icP.Memory[icP.aParam(instruction)] = icP.bParam(instruction) * icP.cParam(instruction)
				icP.Pointer += 4
				return 1
			case 3:
				if icP.Phase == -1 {
					icP.Memory[icP.cParam(instruction)] = icP.Input
				} else {
					if icP.Pointer == 0 {
						icP.Memory[icP.cParam(instruction)] = icP.Phase
					} else {
						icP.Memory[icP.cParam(instruction)] = icP.Input
					}
				}
				icP.Pointer += 2
				return 1
			case 4:
				if icP.DoesRecur {
					icP.Output = append(icP.Output, icP.cParam(instruction))
					icP.Pointer += 2
					return 1
				} else {
					icP.Output = append(icP.Output, icP.cParam(instruction))
					icP.Pointer += 2
					return 0
				}
			case 5:
				if icP.cParam(instruction) != 0 {
					icP.Pointer = icP.bParam(instruction)
				} else {
					icP.Pointer += 3
				}
				return 1
			case 6:
				if icP.cParam(instruction) == 0 {
					icP.Pointer = icP.bParam(instruction)
				} else {
					icP.Pointer += 3
				}
				return 1
			case 7:
				if icP.cParam(instruction) < icP.bParam(instruction) {
					icP.Memory[icP.aParam(instruction)] = 1
				} else {
					icP.Memory[icP.aParam(instruction)] = 0
				}
				icP.Pointer += 4
				return 1
			case 8:
				if icP.cParam(instruction) == icP.bParam(instruction) {
					icP.Memory[icP.aParam(instruction)] = 1
				} else {
					icP.Memory[icP.aParam(instruction)] = 0
				}
				icP.Pointer += 4
				return 1
			case 9:
				icP.RelativeBase += icP.cParam(instruction)
				icP.Pointer += 2
				return 1
			default:
				panic("opcode is not valid")
			}
		}
	}
}
