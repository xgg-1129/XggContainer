package Sort

func SelectSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return selectAscSort(arr)
	case Desc:
		return selectDescSort(arr)
	default:
		return nil
	}
}

func selectAscSort(arr []int) []int {
	n:=len(arr)
	res:=make([]int,n)
	copy(res, arr)
	for i:=0;i<n-1;i++{
		min := i
		for j:=i;j<n;j++{
			//选择排序不稳定
			if res[min]>=res[j]{
				min=j
			}
		}
		res[min],res[i]=res[i],res[min]
	}
	return res
}

func selectDescSort(arr []int) []int {
	n:=len(arr)
	res:=make([]int,n)
	copy(res, arr)
	for i:=0;i<n-1;i++{
		max := i
		for j:=i;j<n;j++{
			if res[max]<=res[j]{
				max=j
			}
		}
		res[max],res[i]=res[i],res[max]
	}
	return res
}
