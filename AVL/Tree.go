package AVL


import "fmt"

/*二叉平衡树是一种基于二叉搜索树的树形数据结构，在AVL树中，他在添加、删除元素时候会更具情况做出旋转的操作，
以此来保证任一节点对应的两棵子树的最大高度差为1，保证了查询效率始终保持再O(logN),它的缺点留到红黑树的时候讲吧
*/




type AVL struct {
 	root *Node
}

func NewAvlTree()*AVL{
	avl:=new(AVL)
	return avl
}
func (tree *AVL) Add(data int64)  {
	tree.root=tree.root.add(data)
}
func (tree *AVL) Get(data int64) *Node {
	return tree.root.get(data)
}
func (tree *AVL) LDR_Print() {
	fmt.Printf("[")
	tree.root.LDR_Print()
	fmt.Println("]")
}
func (tree *AVL) Delete(data int64) {
	if tree.root.get(data)==nil{
		return
	}
	tree.root.delete(data)
}


