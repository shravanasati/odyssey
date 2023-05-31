package parser

import "fmt"

type Node interface {
	String() string
}

type NumberNode struct {
	value string
}

func (nn NumberNode) String() string {
	return nn.value
}

type AddNode struct {
	leftNode  Node
	rightNode Node
}

func (an AddNode) String() string {
	return fmt.Sprintf("(%v + %v)", an.leftNode.String(), an.rightNode.String())
}

type SubstractNode struct {
	leftNode  Node
	rightNode Node
}

func (sn SubstractNode) String() string {
	return fmt.Sprintf("(%v - %v)", sn.leftNode.String(), sn.rightNode.String())
}

type MultiplyNode struct {
	leftNode  Node
	rightNode Node
}

func (mn MultiplyNode) String() string {
	return fmt.Sprintf("(%v * %v)", mn.leftNode.String(), mn.rightNode.String())
}

type DivideNode struct {
	leftNode  Node
	rightNode Node
}

func (dn DivideNode) String() string {
	return fmt.Sprintf("(%v / %v)", dn.leftNode.String(), dn.rightNode.String())
}

type ExponentNode struct {
	leftNode  Node
	rightNode Node
}

func (en ExponentNode) String() string {
	return fmt.Sprintf("(%v ^ %v)", en.leftNode.String(), en.rightNode.String())
}

type ModulusNode struct {
	leftNode  Node
	rightNode Node
}

func (mn ModulusNode) String() string {
	return fmt.Sprintf("(%v %% %v)", mn.leftNode.String(), mn.rightNode.String())
}

type UnaryPlusNode struct {
	node  Node
}

func (un UnaryPlusNode) String() string {
	return fmt.Sprintf("+(%v)", un.node.String())
}

type UnaryMinusNode struct {
	node Node
}

func (un UnaryMinusNode) String() string {
	return fmt.Sprintf("-(%v)", un.node.String())
}