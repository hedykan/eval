package eval

import "strconv"

type operateMap map[string]func(*EvalNode) (any, error)

var opMap operateMap

func init() {
	opMap = make(operateMap)
	opMap["+"] = add
}

func add(en *EvalNode) (any, error) {
	sum := 0
	for _, node := range en.NodeArr[1:] {
		num, err := strconv.Atoi(node.Value)
		if err != nil {
			return nil, err
		}
		sum += num
	}
	return sum, nil
}
