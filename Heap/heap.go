package Heap

type HeapData interface {
	Swap(i,j int)
	Less(i,j int)bool
	Len()int
	Push(interface{})
	Pop() interface{}
}
func Init(data HeapData ){
	//如果左右子树满足堆，则根节点直接和左右子树较大或者较小的交换即可满足堆的定义
	//自下而上，每个根节点下沉
	length:=data.Len()
	for i:=length/2-1;i>=0;i--{
		down(data,i,length)
	}

}
func Push(data HeapData,x interface{}) {
	data.Push(x)
	n:=data.Len()
	up(data,n-1)

}
func Pop(data HeapData)interface{}{
	n:=data.Len()-1
	data.Swap(0,n)
	down(data,0,n)
	return data.Pop()
}
func up(data HeapData, i int) {
	index:=i
	for {
		parent:=(i-1)/2
		if parent<0 || data.Less(index,parent){
			break
		}
		data.Swap(parent,index)
		index=parent
	}
}
func down(data HeapData, i int, length int) {
	index:=i
	for{
		child1:=index*2+1
		//如果没有左孩子，说明自己是叶子结点，down停止
		if child1>=length {
			break
		}
		temp:=child1
		if child2:=child1+1;child2>=length && data.Less(child1,child2){
			temp=child2
		}
		if data.Less(temp,index){
			break
		}
		data.Swap(index,temp)
		index=temp
	}
}




