package day08

type Node struct {
	Id    string
	Left  string
	Right string
}

func (n *Node) IsDeadEnd() bool {
	return n.Id == n.Left && n.Id == n.Right
}
