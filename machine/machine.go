package machine

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func Run(ops []int) []int {
	// Map opcodes to functions
	opCodes := map[int]func(int, int) int{
		1: Add,
		2: Multiply,
	}

	for i := 0; ; i = i + 4 {
		op := ops[i]

		if op == 99 || ops[i+1] > len(ops) || ops[i+2] > len(ops) || ops[i+3] > len(ops) {
			break
		}

		x := ops[ops[i+1]]
		y := ops[ops[i+2]]
		pos := ops[i+3]

		ops[pos] = opCodes[op](x, y)
	}

	return ops
}
