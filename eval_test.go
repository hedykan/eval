package eval

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	node, err := NewEvalNode("[123 [+ 123]]")
	if err != nil {
		t.Error(err.Error())
	}
	if node.Value != "[123 [+ 123]]" {
		t.Error("parse error")
	}
	if node.NodeArr[0].Value != "123" {
		t.Error("parse error")
	}
	if node.NodeArr[1].Value != "[+ 123]" {
		t.Error("parse error")
	}
	if node.NodeArr[1].NodeArr[1].Value != "123" {
		t.Error("parse error")
	}
}

func TestOperate(t *testing.T) {
	node, err := NewEvalNode("[+ 1 [+ 1 2]]")
	if err != nil {
		t.Error(err.Error()) // TODO
	}
	res, err := node.Eval()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(res.(int))
}
