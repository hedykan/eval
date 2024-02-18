package eval

import (
	"errors"
	"strconv"
	"strings"
)

var opMap map[string]func(*EvalNode) (any, error)

func init() {
	opMap = make(map[string]func(*EvalNode) (any, error))
	opMap["+"] = func(en *EvalNode) (any, error) {
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
}

type EvalNode struct {
	Value   string
	NodeArr []*EvalNode
}

func NewEvalNode(str string) (*EvalNode, error) {
	node := new(EvalNode)
	node.Input(str)
	err := node.Parse()
	return node, err
}

func (node *EvalNode) Input(str string) {
	node.Value = str
}

// 解析
func (node *EvalNode) Parse() error {
	src, err := getValue(node.Value) // 是否可以解析，否则将返回err
	if err != nil {
		return err
	}
	valueArr := divValue(src)
	for _, value := range valueArr {
		subNode, err := NewEvalNode(value)
		if err != nil {
			return err
		}
		node.NodeArr = append(node.NodeArr, subNode)
	}
	return nil
}

var left = '['
var right = ']'

// 获取括号内容
func getValue(str string) (string, error) {
	// 是否存在表达式
	findStatus := false
	find := 0
	startIndex := 0
	endIndex := 0
	for i := 0; i < len(str); i++ {
		if str[i] == byte(left) {
			if find == 0 {
				startIndex = i
			}
			find += 1
		}
		if str[i] == byte(right) {
			find -= 1
			if find == 0 {
				endIndex = i
				findStatus = true
			}
		}
	}
	if findStatus {
		return str[startIndex+1 : endIndex], nil
	} else {
		var err error
		if find != 0 {
			err = errors.New("get value error")
		} else {
			err = nil
		}
		return "", err
	}
}

// 分割值
func divValue(str string) []string {
	return strings.Fields(str)
}
