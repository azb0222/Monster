package parser

import (
	"fmt"
	"strconv"
)

type Expression interface {
}

// Actual Types
type Literal struct {
}

// interface for evaluatable expressions
type Eval interface {
	Eval() float64
}

type Stringify interface {
	ToString() string
}

// Types
type Constant struct {
	value float64
}

type BinaryPlus struct {
	left  Expression //can have arbritary types at compile time
	right Expression
}

func (c *Constant) Eval() float64 {
	return c.value
}

func (bp *BinaryPlus) Eval() float64 {
	return bp.left.(Eval).Eval() + bp.right.(Eval).Eval() //type casting to Eval
}

func (c *Constant) ToString() string {
	return strconv.FormatFloat(c.value, 'f', -1, 64)
}

func (bp *BinaryPlus) ToString() string {
	ls := bp.left.(Stringify).ToString()
	rs := bp.right.(Stringify).ToString()
	return fmt.Sprintf("%s %s", ls, rs)
}
