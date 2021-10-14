package BinarySortTree

import (
	"testing"
)

func Test(t *testing.T) {
	values := []int{3, 6, 8, 20, 9, 2, 6, 8, 9, 3, 5, 40, 7, 9, 13, 6, 8}
	tree:=NewTree()
	for _,value := range values{
		tree.Add(value)
	}
	//中序打印
	tree.LDR_Print()

	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		t.Error("find is fault")
	}
	// 查找存在的9
	node = tree.Find(9)
	if node == nil {
		t.Error("can not find value")
	}

	// 删除存在的9后，再查找9
	tree.Delete(9)
	node = tree.Find(9)
	if node != nil {
		t.Error("delete faule")
	}

	tree.LDR_Print()
}
