package machine

import "fmt"

const (
	_ = iota
	opAdd
	opMultiply
	opInput
	opOutput
	opJumpT
	opJumpF
	opLess
	opEqual
	opAdjustRel

	opHalt = 99
)

type Machine struct {
	IP           int
	RelativeBase int
	Operation    int
	Instruction  int
	Output       int
	Memory       []int
	MockedInput  []int
	MockIndex    int
	StoreOutput  bool
	Mock         bool
	Halted       bool
	RunSetTimes  bool
	RunCounter   int
}

func New(opcodes []int) *Machine {
	memory := make([]int, len(opcodes)*100)
	copy(memory, opcodes)

	return &Machine{
		Memory: memory,
	}
}

// Parameter Modes
//
// 0 - Position    m.Memory[m.Memory[m.IP]]
// 1 - Immediate   m.Memory[m.IP]
// 2 - Relative    m.Memory[m.RelativeBase+m.Memory[m.IP]]
func (m *Machine) oneParamater() int {
	paramModes := m.Instruction - m.Operation

	pm1 := paramModes % 1000 / 100

	m.IP++
	p1 := m.Memory[m.IP]
	if pm1 == 0 {
		p1 = m.Memory[p1]
	}
	if pm1 == 2 {
		if m.Operation == opInput {
			p1 = m.RelativeBase + p1
		} else {
			p1 = m.Memory[m.RelativeBase+p1]
		}
	}

	return p1
}

func (m *Machine) twoParameters() (int, int) {
	paramModes := m.Instruction - m.Operation

	p1 := m.oneParamater()
	pm2 := paramModes % 10000 / 1000

	m.IP++
	p2 := m.Memory[m.IP]
	if pm2 == 0 {
		p2 = m.Memory[p2]
	}
	if pm2 == 2 {
		p2 = m.Memory[m.RelativeBase+p2]
	}

	return p1, p2
}

func (m *Machine) threeParameters() (int, int, int) {
	paramModes := m.Instruction - m.Operation

	p1, p2 := m.twoParameters()
	pm3 := paramModes % 100000 / 10000

	m.IP++
	p3 := m.IP
	if pm3 == 0 {
		p3 = m.Memory[p3]
	}
	if pm3 == 2 {
		p3 = m.RelativeBase + p3
	}

	return p1, p2, p3
}

func (m *Machine) parseInstruction() {
	m.Instruction = m.Memory[m.IP]
	m.Operation = m.Instruction % 100
}

func (m *Machine) halt() {
	m.Halted = true
}

func (m *Machine) add() {
	p1, p2, p3 := m.threeParameters()

	m.Memory[p3] = p1 + p2
}

func (m *Machine) multiply() {
	p1, p2, p3 := m.threeParameters()

	m.Memory[p3] = p1 * p2
}

func (m *Machine) input() {
	var v int
	if !m.Mock {
		_, err := fmt.Scanf("%d", &v)
		if err != nil {
			panic("input error")
		}
	} else {
		v = m.MockedInput[m.MockIndex]
		m.MockIndex++
	}

	p1 := m.oneParamater()
	m.Memory[p1] = v
}

func (m *Machine) output() {
	p1 := m.oneParamater()
	if m.StoreOutput {
		m.Output = p1
	} else {
		fmt.Println(p1)
	}
}

func (m *Machine) jumpIfTrue() {
	p1, p2 := m.twoParameters()

	if p1 != 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) jumpIfFalse() {
	p1, p2 := m.twoParameters()

	if p1 == 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) lessThan() {
	p1, p2, p3 := m.threeParameters()

	if p1 < p2 {
		m.Memory[p3] = 1
	} else {
		m.Memory[p3] = 0
	}
}

func (m *Machine) equals() {
	p1, p2, p3 := m.threeParameters()

	if p1 == p2 {
		m.Memory[p3] = 1
	} else {
		m.Memory[p3] = 0
	}
}

func (m *Machine) adjustRelativeBase() {
	p1 := m.oneParamater()
	m.RelativeBase += p1
}

func (m *Machine) DumpMemory() []int {
	return m.Memory
}

func (m *Machine) MockInput(input []int) {
	m.Mock = true

	for _, v := range input {
		m.MockedInput = append(m.MockedInput, v)
	}
}

func (m *Machine) SaveOutput() {
	m.StoreOutput = true
}

func (m *Machine) RunFor(times int) {
	m.RunSetTimes = true
	m.RunCounter = times
}

func (m *Machine) Run() {
	for !m.Halted {
		m.parseInstruction()

		//fmt.Println(m.IP, m.Instruction)

		switch m.Operation {
		case opHalt:
			m.halt()
			return
		case opAdd:
			m.add()
		case opMultiply:
			m.multiply()
		case opInput:
			m.input()
		case opOutput:
			m.output()
			if m.RunSetTimes {
				m.RunCounter--
				if m.RunCounter == 0 {
					m.RunSetTimes = false
					m.IP++
					return
				}
			}

		case opJumpT:
			m.jumpIfTrue()
		case opJumpF:
			m.jumpIfFalse()
		case opLess:
			m.lessThan()
		case opEqual:
			m.equals()
		case opAdjustRel:
			m.adjustRelativeBase()
		default:
			panic("unrecognized opcode")
		}
		m.IP++
	}
}
