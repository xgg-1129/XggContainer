package Sort

/*计数排序一般不适合有负数的情况，但是也可以硬写一个，比如这个,但是这不适用于数组里有-217..和217..的情况*/
func CountSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return countAscSort(arr)
	case Desc:
		return countDescSort(arr)
	default:
		return nil
	}
}



func countAscSort(arr []int) []int {
	n:=len(arr)
	var res []int
	if n==0||n==1{
		return  res
	}
	//求出最大最小值
	min,max := arr[0],arr[0]
	for _,item := range arr{
		if min>item{
			min=item
		}
		if max<item{
			max=item
		}
	}
	//构建计数数组
	length:=max-min+1
	countDp:=make([]int,length)
	//开始计数
	for _,item :=range arr{
		countDp[item-min]++
	}

	//填充结果数组
	for i,v := range  countDp{
		for j:=0;j<v;j++{
			res=append(res,i+min)
		}
	}
	return res
}

func countDescSort(arr []int) []int {
	res:=countAscSort(arr)
	Reverse(res)
	return res
}