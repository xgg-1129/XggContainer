package AVL

import "fmt"

type Node struct {
	data int64
	left *Node
	right *Node
	time int64

	//以该节点为root的高度
	Height int64

}
//获取高度
func (n *Node) GetHeight()int64{
	if n==nil{
		return 0
	}
	leftHeight:=n.left.GetHeight()
	rightHeight:=n.right.GetHeight()
	if leftHeight>rightHeight{
		n.Height=leftHeight+1
	}else{
		n.Height=rightHeight+1
	}
	return n.Height
}

//计算平衡因子
func (n *Node) Factory() int64 {
	return n.left.GetHeight()-n.right.GetHeight()
}

//左旋操作
func Lswirl(oldRoot *Node) *Node{
	pivot:=oldRoot.right
	oldRoot.right=pivot.left
	pivot.left=oldRoot

	pivot.GetHeight()
	oldRoot.GetHeight()
	return pivot
}
//右旋操作
func Rswirl(oldRoot *Node)*Node{
	pivot:=oldRoot.left
	oldRoot.left=pivot.right
	pivot.right=oldRoot

	pivot.GetHeight()
	oldRoot.GetHeight()
	return pivot
}
//先左旋后右旋
func LRswirl(oldRoot *Node)*Node{
	oldRoot.left= Lswirl(oldRoot.left)
	return Rswirl(oldRoot)
}

//先右旋再左旋
func RLswirl(oldRoot *Node)*Node{
	oldRoot.right= Rswirl(oldRoot.right)
	return Lswirl(oldRoot)
}

func (root *Node) add(data int64) (*Node){
	if root==nil{
		return &Node{
			data:   data,
			time:   0,
			Height: 1,
		}
	}
	if data==root.data{
		root.time++
		return root
	}

	var newRootNode *Node
	if root.data<data{
		root.right=root.right.add(data)
		factory := root.Factory()
		//如果该二叉树失衡
		if factory==-2{
				//左旋
				if data>root.right.data{
					newRootNode =Lswirl(root)
				}else{
					newRootNode = RLswirl(root)
				}
		}
	}else{
		root.left=root.left.add(data)
		factory := root.Factory()
		if factory == 2{
			if data<root.left.data{
				newRootNode=Rswirl(root)
			}else{
				newRootNode=LRswirl(root)
			}
		}
	}
	if newRootNode==nil{
		root.GetHeight()
		return root
	}else{
		newRootNode.GetHeight()
		return newRootNode
	}
}

func (root *Node) get(data int64)*Node{
	if root==nil{
		return root
	}

	if root.data==data{
		return root
	}else if root.data < data{
		return root.right.get(data)
	} else{
		return root.left.get(data)
	}
}

func (root *Node) LDR_Print() {
	if root==nil{
		return
	}
	root.left.LDR_Print()
	for i:=int64(-1);i < root.time;i++{
		fmt.Printf(" %d ",root.data)
	}
	root.right.LDR_Print()
}
func (root *Node) GetMax()*Node{
	if root==nil{
		return nil
	}
	if root.right == nil{
		return root
	}
	return root.right.GetMax()
}
func (root *Node) GetMin() *Node{
	if root == nil{
		return nil
	}
	if root.left==nil{
		return root
	}
	return root.left.GetMin()
}

func (root *Node) delete(data int64)*Node{
	if root == nil{
		return nil
	}
	if root.data<data{
		root.right=root.right.delete(data)
	}else if root.data>data{
		root.left=root.left.delete(data)
	}else{  //删除自己

		//1.如果自己是叶子节点
		if root.left==nil&&root.right==nil{
			return nil
		}

		//如果只有左子树

		if root.left!=nil&&root.right==nil{
			root.data=root.left.data
			root.time=root.left.time
			root.Height=root.left.Height
			root.left=nil
			return root
		}
		//只有右子树
		if root.left==nil&&root.right!=nil{
			root.data=root.right.data
			root.time=root.right.time
			root.Height=root.right.Height
			root.right=nil
			return root
		}

		/*删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，如果左子树
		更高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，如果右子树更高，选择右子树中
		最小的节点，也就是右子树最左边的叶子节点。最后，删除这个叶子节点
		 */
		if root.left.Height>root.right.Height{
			node:=root.left.GetMax()
			root.data=node.data
			root.time=node.time
			root.left.delete(node.data)
			root.GetHeight()
			return root
		}

		if root.left.Height<root.right.Height{
			node:=root.right.GetMin()
			root.data=node.data
			root.time=node.time
			root.right.delete(node.data)
			root.GetHeight()
			return root
		}
	}
	root.GetHeight()
	var newnode *Node
	if root.Factory()==2{
		if root.left.Factory() >=0{
			newnode=Rswirl(root)
		}else{
			newnode=LRswirl(root)
		}
	}else if root.Factory()==-2{
		if root.right.Factory()<=0{
			newnode=Lswirl(root)
		}else{
			newnode=RLswirl(root)
		}
	}

	if newnode==nil{
		return root
	}else{
		newnode.GetHeight()
		return newnode
	}
}


