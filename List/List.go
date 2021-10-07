package List

type Object interface {}


type linkedlistnode struct {
	data Object
	next *linkedlistnode
}

type LinkedList struct {
	head *linkedlistnode
}





