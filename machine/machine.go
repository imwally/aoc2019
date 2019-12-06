package machine

import "fmt"

func Run(ops []int) []int {
	for i := 0; ; i++ {
		ins := ops[i]
		op := ins % 100

		if op == 99 {
			break
		}

		p1, p2 := 0, 0
		params := ins - op

		pm1 := params % 1000 / 100
		pm2 := params % 10000 / 1000

		i++
		if pm1 == 0 {
			p1 = ops[ops[i]]
		} else {
			p1 = ops[i]
		}

		if op < 3 || op == 5 || op == 6 || op == 7 || op == 8 {
			i++
			if pm2 == 0 {
				p2 = ops[ops[i]]
			} else {
				p2 = ops[i]
			}
		}

		switch op {
		case 1:
			i++ // Parameter 3
			ops[ops[i]] = p1 + p2
		case 2:
			i++ // Parameter 3
			ops[ops[i]] = p1 * p2
		case 3:
			var v int
			_, err := fmt.Scanf("%d", &v)
			if err != nil {
				panic("shit broke")
			}
			ops[ops[i]] = v
		case 4:
			fmt.Println(p1)
		case 5:
			if p1 != 0 {
				i = p2 - 1
			}
		case 6:
			if p1 == 0 {
				i = p2 - 1
			}
		case 7:
			i++ // Parameter 3
			if p1 < p2 {
				ops[ops[i]] = 1
			} else {
				ops[ops[i]] = 0
			}
		case 8:
			i++ // Parameter 3
			if p1 == p2 {
				ops[ops[i]] = 1
			} else {
				ops[ops[i]] = 0
			}
		default:
			panic("unrecognized opcode")
		}

	}

	return ops
}
