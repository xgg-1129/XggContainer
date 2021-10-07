package Stack

import (

	"fmt"
)

func GetStack()stack{
	return stack{
		arr: make([]Object,0),
	}
}

func (this *stack) Push(data Object)  {
	this.arr=append(this.arr,data)
}


func (this *stack) Pop()bool  {
	if this.IsEmpty()==true{
		return false
	}else{
		this.arr=this.arr[0:len(this.arr)-1]
		return true
	}
}


func (this *stack) IsEmpty() bool {
	if n:=len(this.arr);n==0{
		return true
	}
	return false
}


func (this *stack) Top() Object {
	return this.arr[len(this.arr)-1]
}


func (this *stack) PrintStack()  {
	n:=len(this.arr)
	for i:=0;i<n;i++{
		fmt.Println(this.arr[i])
	}
}


