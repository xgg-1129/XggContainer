package Slice

import (
	"fmt"
	"sync"
)

type Slice struct {
	Array []interface{}
	len int
	cap int

	mu sync.Mutex
}

func (s *Slice) Cap() int {
	return s.cap
}

func (s *Slice) Len() int {
	return s.len
}

func NewSlice(len,cap int)*Slice{
	if len>cap{
		panic("len larger than cap")
	}
	res:=new(Slice)
	res.Array=make([]interface{},cap,cap)
	res.len=len
	res.cap=cap
	return res
}

func Append(slice *Slice,element interface{})*Slice{
	slice.mu.Lock()
	defer slice.mu.Unlock()
	//判断是否需要扩容
	if slice.len==slice.cap{
		newcap:=slice.cap*2
		newArray:=make([]interface{},newcap,newcap)
		for index,value := range slice.Array{
			newArray[index]=value
		}
		slice.Array=newArray
		slice.cap=newcap
	}
	slice.Array[slice.len]=element
	slice.len++
	return slice
}

func (s *Slice) Get(index int)interface{}{
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.len==0|| index>=s.len {
		panic("over slice len")
	}
	return s.Array[index]
}

func (s *Slice) PrintSlice(){
	fmt.Print("[")
	for i := 0; i < s.Len(); i++ {
		fmt.Print(s.Array[i])
		if i==s.Len()-1{
			break
		}
		fmt.Print(",")
	}
	fmt.Print("]\n")
}
