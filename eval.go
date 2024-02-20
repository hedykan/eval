package eval

import (
	"errors"
	"strings"
)

type EvalNode struct {
	Value     string
	NodeArr   []*EvalNode
	ValueType int // 0:symbol,1:int,2:string
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

// 求值，循环往下求
func (node *EvalNode) Eval() (any, error) {
	if len(node.NodeArr) == 0 {
		return node.Value, nil
	}
	op, err := node.NodeArr[0].Eval()
	if err != nil {
		return nil, err
	}
	detailOp := op.(string)
	if foo, ok := opMap[detailOp]; !ok {
		return nil, errors.New("function not found")
	} else {
		res, err := foo(node)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
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

// TODO 跳过""号
// 分割值
func divValue(str string) []string {
	return strings.Fields(str)
}
