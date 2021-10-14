package BinarySortTree

import (
	"fmt"
)


/*
	二叉查找树是一种基于二分查找的数据结构，它的时间复杂度在于树的高度，O(LogN)
	在极端的情况下，二叉查找树会退化为链表，所以工程上一般会使用AVL和红黑树，它们
	有维持自生高度的平衡操作，能够始终保持查找效率为O(LogN)
*/

type BinarySortTree struct {
	root *Node
}
type do func(*Node)
type Node struct {
	data int
	time int
	left *Node
	right *Node
}


func (root *Node) add(data int) {
	if root.data == data{
		root.time++
		return
	}else if root.data<data{
		if root.right==nil{
			root.right=&Node{data: data}
			return
		}
		root.right.add(data)
	}else{
		 if root.left==nil{
		 	root.left=&Node{data: data}
		 	return
		 }
		 root.left.add(data)
	}
}
func NewBinarySortTree()*BinarySortTree{
	tree:=new(BinarySortTree)
	return tree
}
func (t *BinarySortTree) Add(data int) {
	 if t.root==nil{
	 	t.root=&Node{data: data}
		 return
	 }
	 t.root.add(data)
}
func (root *Node) LDR(f do){
	if  root==nil{
		return
	}
	root.left.LDR(f)
	f(root)
	root.right.LDR(f)
}

func (t *BinarySortTree) LDR_Print()  {
	t.root.ldr_Print()
}

func (root *Node) ldr_Print(){
	if root== nil{
		return
	}
	root.left.ldr_Print()
	fmt.Println(root.data)
	root.right.ldr_Print()
}
func (root *Node) find(value int)*Node{
	if root == nil{
		return nil
	}else if root.data==value{
		return root
	}else if root.data<value{
		return root.right.find(value)
	}else{
		return root.left.find(value)
	}
}
func (t *BinarySortTree) Find(value int)*Node{
	return t.root.find(value)
}
func (root *Node) findParent(value int) *Node{
	if root== nil || root.data==value{
		return nil
	}else if root.data<value{
		if root.right==nil{
			return nil
		}
		if root.right.data==value{
			return root
		}else{
			return root.right.findParent(value)
		}
	}else{
		if root.data > value{
			if root.left==nil{
				return nil
			}
			if root.left.data==value{
				return root
			}else{
				return root.left.findParent(value)
			}
		}
	}
	return nil
}
//删除需要考虑4种情况
func (t *BinarySortTree) Delete(data int) {
	//树为空则直接return
	if t.root==nil{
		return
	}
	//如果节点不存在则直接返回
	node:=t.Find(data)
	if node==nil{
		return
	}
	parent:=t.root.findParent(data)
	//如果删除的是根结点，且没有子节点
	if node == t.root && node.left==nil && node.right==nil{
		t.root=nil
		return

	//如果删除的是叶子结点
	}else if node.left == nil && node.right == nil{
		if parent.left==node{
			parent.left=nil
		}else{
			parent.right=nil
		}
		return
	//删除的结点有左右字数,用右子树的最小值取代自己
	}else if node.left!=nil && node.right!=nil{
		minNode:=node.right
		for minNode.left==nil{
			minNode=minNode.left
		}
		t.Delete(minNode.data)
		node.data=minNode.data
		node.time=minNode.time
		return
	//node结点只有一个子树
	}else{
		//如果是根结点
		if node==t.root{
			if t.root.left!=nil{
				t.root=t.root.left
				//断开根节点和树的连接
				node.left=nil
			}else{
				t.root=t.root.right
				//同上
				node.right=nil
			}
		}else if node.left!=nil{
			if parent.left==node{
				parent.left=node.left
			}else{
				parent.right=node.left
			}
			node.left=nil
		}else{
			if parent.left==node{
				parent.left=node.right
			}else{
				parent.right=node.right
			}
			node.right=nil
		}
	}

}