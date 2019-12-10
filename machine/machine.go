package machine

import "fmt"

type Machine struct {
	IP         int
	Operation  int
	Parameters []int
	Memory     []int
	Halted     bool
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
		Halted:     false,
	}
}

func (m *Machine) parseInstruction() {
	ins := m.Memory[m.IP]
	m.Operation = ins % 100

	// Don't parse parameter if machine needs to halt
	if m.Operation == 99 {
		m.Halted = true
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

	// Only certain operations such as Addition, Mutliplication, and
	// Jumping require a second parameter.
	if m.Operation == 5 || m.Operation == 6 {
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

func (m *Machine) Add() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	m.Memory[p3] = p1 + p2
}

func (m *Machine) Multiply() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[1]

	m.IP++
	p3 := m.Memory[m.IP]
	m.Memory[p3] = p1 * p2
}

func (m *Machine) Input() {
	var v int
	_, err := fmt.Scanf("%d", &v)
	if err != nil {
		panic("shit broke")
	}
	m.IP++
	p3 := m.Memory[m.IP]
	m.Memory[p3] = v
}

func (m *Machine) Output() {
	p1 := m.Parameters[0]

	fmt.Println(p1)
}

func (m *Machine) JumpIfTrue() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[2]

	if p1 != 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) JumpIfFalse() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[2]

	if p1 == 0 {
		m.IP = p2 - 1
	}
}

func (m *Machine) LessThan() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[2]

	m.IP++
	p3 := m.Memory[m.IP]
	if p1 < p2 {
		m.Memory[p3] = 1
	} else {
		m.Memory[p3] = 0
	}
}

func (m *Machine) Equals() {
	p1 := m.Parameters[0]
	p2 := m.Parameters[2]

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

func (m *Machine) Run() {
	for !m.Halted {
		m.parseInstruction()

		switch m.Operation {
		case 99:
			m.Halted = true
			return
		case 1:
			m.Add()
		case 2:
			m.Multiply()
		case 3:
			m.Input()
		case 4:
			m.Output()
		case 5:
			m.JumpIfTrue()
		case 6:
			m.JumpIfFalse()
		case 7:
			m.LessThan()
		case 8:
			m.Equals()
		default:
			panic("unrecognized opcode")
		}
		m.IP++
	}
}
