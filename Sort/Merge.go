package Sort


func MergeSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return mergeAscSort(arr)
	case Desc:
		return mergeDescSort(arr)
	default:
		return nil
	}
}

func mergeDescSort(arr []int) []int {
	res := mergeAscSort(arr)
	Reverse(res)
	return res
}

func mergeAscSort(arr []int)[]int{
	n:=len(arr)
	if n==0||n==1{
		return arr
	}
	mid:=n/2
	a:=mergeAscSort(arr[0:mid])
	b:=mergeAscSort(arr[mid:])
	return merge(a,b)
}


func merge(a,b []int) []int{
	var res []int
	aindex,bindex:=0,0
	for aindex<len(a)&&bindex<len(b){
		if a[aindex]>b[bindex]{
			res=append(res,b[bindex])
			bindex++
		}else{
			res=append(res,a[aindex])
			aindex++
		}
	}
	if aindex>=len(a){
		res=append(res,b[bindex:]...)
	}else{
		res=append(res,a[aindex:]...)
	}
	return res
}
