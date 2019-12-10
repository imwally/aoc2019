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

	opHalt = 99
)

type Machine struct {
	IP          int
	Operation   int
	Parameters  []int
	Memory      []int
	MockedInput int
	Mock        bool
	Halted      bool
}

func New(opcodes []int) *Machine {
	parameters := make([]int, 3)
	memory := make([]int, len(opcodes))
	copy(memory, opcodes)
	return &Machine{
		IP:         0,
		Operation:  0,
		Parameters: parameters,
		Memory:     memory,
		Mock:       false,
		Halted:     false,
	}
}

func (m *Machine) parseInstruction() {
	ins := m.Memory[m.IP]
	m.Operation = ins % 100

	// Don't parse parameter if machine needs to halt
	if m.Operation == opHalt {
		m.halt()
		return
	}

	// Parameter Modes
	//
	// Assume immediate mode by default. This allows us easily update the
	// parameter to position mode if need be.
	//
	// 1 - Immediate Mode	p1 := m.Memory[m.IP]	Value at IP
	// 0 - Position Mode	p1 = m.Memory[p1]	Value at what IP points to
	paramModes := ins - m.Operation
	pm1 := paramModes % 1000 / 100
	pm2 := paramModes % 10000 / 1000

	// First Parameter
	//
	// All operations assume there is at least one parameter. This is the
	// next value in memory.
	m.IP++
	m.Parameters[0] = m.Memory[m.IP]
	if pm1 == 0 {
		m.Parameters[0] = m.Memory[m.Parameters[0]]
	}

	// Certain operations only require a single parameter. Don't parse or
	// increase the IP for those operations.
	if m.Operation == opInput || m.Operation == opOutput {
		// Stop further parsing
		return
	}

	// Second Parameter
	//
	// Next value in memory after the first parameter
	m.IP++
	m.Parameters[1] = m.Memory[m.IP]
	if pm2 == 0 {
		m.Parameters[1] = m.Memory[m.Parameters[1]]
	}
}

func (m *Machine) halt() {
	m.Halted = true
}

func (m *Machine) add() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	m.Memory[p3] = p1 + p2
}

func (m *Machine) multiply() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	m.Memory[p3] = p1 * p2
}

func (m *Machine) input() {
	var v int
	if !m.Mock {
		_, err := fmt.Scanf("%d", &v)
		if err != nil {
			panic("shit broke")
		}
	} else {
		v = m.MockedInput
	}

	m.Memory[m.Memory[m.IP]] = v
}

func (m *Machine) output() {
	p1 := m.Parameters[0]
	fmt.Println(p1)
}

func (m *Machine) jumpIfTrue() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	if p1 != 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) jumpIfFalse() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	if p1 == 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) lessThan() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	if p1 < p2 {
		m.Memory[p3] = 1
	} else {
		m.Memory[p3] = 0
	}
}

func (m *Machine) equals() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	if p1 == p2 {
		m.Memory[p3] = 1
	} else {
		m.Memory[p3] = 0
	}
}

func (m *Machine) DumpMemory() []int {
	return m.Memory
}

func (m *Machine) MockInput(v int) {
	m.Mock = true
	m.MockedInput = v
}

func (m *Machine) Run() {
	for !m.Halted {
		m.parseInstruction()

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
		case opJumpT:
			m.jumpIfTrue()
		case opJumpF:
			m.jumpIfFalse()
		case opLess:
			m.lessThan()
		case opEqual:
			m.equals()
		default:
			panic("unrecognized opcode")
		}
		m.IP++
	}
}
