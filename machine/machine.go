package machine

import "fmt"

func Run(ops []int) []int {
	for i := 0; ; {
		ins := ops[i]
		op := ins % 100

		if op == 99 {
			break
		}

		p1, p2, p3 := 0, 0, 0
		if ins > 1000 {
			params := ins - op

			pm1 := params % 100
			pm2 := params % 10000 / 1000
			pm3 := params % 100000 / 10000

			i++
			if pm1 == 0 {
				p1 = ops[ops[i]]
			} else {
				p1 = ops[i]
			}

			i++
			if pm2 == 0 {
				p2 = ops[ops[i]]
			} else {
				p2 = ops[i]
			}

			i++
			if pm3 == 0 {
				p3 = ops[i]
			} else {
				p3 = i
			}
			i++
		} else {
			i++
			p1 = ops[ops[i]]
			i++
			p2 = ops[ops[i]]
			i++
			p3 = ops[i]
			i++
		}

		switch op {
		case 1:
			ops[p3] = p1 + p2
		case 2:
			ops[p3] = p1 * p2
		case 3:
			var v int
			_, err := fmt.Scanf("%d", &v)
			if err != nil {
				panic("shit broke")
			}
			ops[p3] = v
		case 4:
			fmt.Println(p1)
		}

	}

	return ops
}
