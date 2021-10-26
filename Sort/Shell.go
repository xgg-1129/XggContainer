package Sort


func ShellSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return shellAscSort(arr)
	case Desc:
		return shellDescSort(arr)
	default:
		return nil
	}
}
//可以用二分查找改进
func shellAscSort(arr []int) []int {
	n:=len(arr)
	res:=make([]int,n)
	copy(res,arr)
	for i:=n/2;i>0;i=i/2{
		for j:=i;j<n;j++{
			index:=j
			for index-i>=0 && res[index]<res[index-i]{
				res[index],res[index-i]=res[index-i],res[index]
				index=index-i
			}
		}
	}
	return res
}

func shellDescSort(arr []int) []int {
	res := shellAscSort(arr)
	Reverse(res)
	return res
}
