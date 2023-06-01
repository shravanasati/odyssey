package interpreter

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Shravan-1908/odyssey/parser"
)

var errZeroDivision = errors.New("zero division error")

func stringToFloat(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)
	if e != nil {
		panic(fmt.Sprintf("stringToFloat(%s): not convertable", s))
	}
	return f
}

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (ip *Interpreter) Visit(node parser.Node) float64 {
	switch node.(type) {
	case parser.NumberNode:
		return ip.visitNumberNode(node.(parser.NumberNode))
	case parser.AddNode:
		return ip.visitAddNode(node.(parser.AddNode))
	case parser.SubstractNode:
		return ip.visitSubstractNode(node.(parser.SubstractNode))
	case parser.MultiplyNode:
		return ip.visitMultiplyNode(node.(parser.MultiplyNode))
	case parser.DivideNode:
		return ip.visitDivideNode(node.(parser.DivideNode))
	case parser.ExponentNode:
		return ip.visitExponentNode(node.(parser.ExponentNode))
	case parser.ModulusNode:
		return ip.visitModulusNode(node.(parser.ModulusNode))
	default:
		panic(fmt.Sprintf("unknown node: %v \n", node))
	}
}

func (ip *Interpreter) visitNumberNode(node parser.NumberNode) float64 {
	return stringToFloat(node.Value)
}

func (ip *Interpreter) visitAddNode(node parser.AddNode) float64 {
	val := ip.Visit(node.LeftNode) + ip.Visit(node.RightNode)
	return val
}

func (ip *Interpreter) visitSubstractNode(node parser.SubstractNode) float64 {
	val := ip.Visit(node.LeftNode) - ip.Visit(node.RightNode)
	return val
}

func (ip *Interpreter) visitMultiplyNode(node parser.MultiplyNode) float64 {
	val := ip.Visit(node.LeftNode) * ip.Visit(node.RightNode)
	return val
}

func (ip *Interpreter) visitDivideNode(node parser.DivideNode) float64 {
	a := ip.Visit(node.LeftNode)
	b := ip.Visit(node.RightNode)
	if b == float64(0) {
		fmt.Println(errZeroDivision.Error())
		return 0
	}
	val := a / b
	return val
}

func (ip *Interpreter) visitExponentNode(node parser.ExponentNode) float64 {
	val := math.Pow(ip.Visit(node.LeftNode), ip.Visit(node.RightNode))
	return val
}

func (ip *Interpreter) visitModulusNode(node parser.ModulusNode) float64 {
	val := int(ip.Visit(node.LeftNode)) % int(ip.Visit(node.RightNode))
	return float64(val)
}
