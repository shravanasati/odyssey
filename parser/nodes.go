package parser

import "fmt"

type Node interface {
	String() string
}

type NumberNode struct {
	Value string
}

func (nn NumberNode) String() string {
	return nn.Value
}

type AddNode struct {
	LeftNode  Node
	RightNode Node
}

func (an AddNode) String() string {
	return fmt.Sprintf("(%v + %v)", an.LeftNode.String(), an.RightNode.String())
}

type SubstractNode struct {
	LeftNode  Node
	RightNode Node
}

func (sn SubstractNode) String() string {
	return fmt.Sprintf("(%v - %v)", sn.LeftNode.String(), sn.RightNode.String())
}

type MultiplyNode struct {
	LeftNode  Node
	RightNode Node
}

func (mn MultiplyNode) String() string {
	return fmt.Sprintf("(%v * %v)", mn.LeftNode.String(), mn.RightNode.String())
}

type DivideNode struct {
	LeftNode  Node
	RightNode Node
}

func (dn DivideNode) String() string {
	return fmt.Sprintf("(%v / %v)", dn.LeftNode.String(), dn.RightNode.String())
}

type ExponentNode struct {
	LeftNode  Node
	RightNode Node
}

func (en ExponentNode) String() string {
	return fmt.Sprintf("(%v ^ %v)", en.LeftNode.String(), en.RightNode.String())
}

type ModulusNode struct {
	LeftNode  Node
	RightNode Node
}

func (mn ModulusNode) String() string {
	return fmt.Sprintf("(%v %% %v)", mn.LeftNode.String(), mn.RightNode.String())
}

type UnaryPlusNode struct {
	node Node
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
