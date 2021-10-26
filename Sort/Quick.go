package Sort

func QuickSort(arr []int,module Type)[]int{
	n:=len(arr)
	res:=make([]int,n)
	copy(res,arr)
	switch module {
	case Asc:
		quickAscSort(res,0,len(res)-1)
	case Desc:
		quickDescSort(res,0,len(res)-1)
	default:
		return nil
	}
	return res
}
func quickAscSort(arr []int,left int,right int){
	if left>=right{
		return
	}
	index:=partiton(arr,left,right)
	quickAscSort(arr,left,index-1)
	quickAscSort(arr,index+1,right)
}
func partiton(arr[]int,left int,right int)(index int){
	index =left
	pivot:=left+1
	for i:=pivot;i<=right;i++{
		if arr[i]<arr[index]{
			arr[i],arr[pivot]=arr[pivot],arr[i]
			pivot++
		}
	}
	arr[index],arr[pivot-1]=arr[pivot-1],arr[index]
	index = pivot-1
	return
}
func quickDescSort(arr []int,left int,right int){
	quickAscSort(arr,left,right)
	Reverse(arr)
}


