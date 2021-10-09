package HashTable


//基于拉链法解决，结点设计如下
type Node struct {
	key string
	value interface{}
	next *Node
}
