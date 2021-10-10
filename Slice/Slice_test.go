package Slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 创建一个容量为3的动态数组
	a := NewSlice(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len())
	a.PrintSlice()
	// 增加一个元素
	Append(a,10)
	fmt.Println("cap", a.Cap(), "len", a.Len())
	a.PrintSlice()
	// 增加一个元素
	Append(a,9)
	fmt.Println("cap", a.Cap(), "len", a.Len())
	a.PrintSlice()
	Append(a,7)
	Append(a,8)
	a.PrintSlice()

}