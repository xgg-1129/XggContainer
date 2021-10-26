package Sort


/*不能处理有负数的情况，有空写个改进版的*/

func BucketSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return bucketAscSort(arr)
	case Desc:
		return bucketDescSort(arr)
	default:
		return nil
	}
}

func bucketDescSort(arr []int) []int {
	res := bucketAscSort(arr)
	Reverse(res)
	return res
}


//通的数量取决于max
func bucketAscSort(arr []int) []int {
	n:=len(arr)
	var res []int
	if n==0||n==1{
		return  res
	}
	//求出最大值
	max := arr[0]
	for _,item := range arr{
		if max<item{
			max=item
		}
	}
	//求出桶的数量
	BucketNum:= max/10 +1
	Buckets:=make([][]int,BucketNum)
	//为每个元素分配桶
	for _,item := range arr{
		index:=item/10
		Buckets[index]=append(Buckets[index],item)
	}
	//给每个桶排序
	for _,b := range Buckets{
		sortOneBucker(b)
	}
	//合并结果
	for _,b:=range Buckets{
		for _,v := range b{
			res=append(res,v)
		}
	}
	return  res
}
func sortOneBucker(res[]int){
	n:=len(res)
	for i:=1;i<n;i++{
		index :=i
		//如果比前面的数小
		for index>0&& res[index] < res[index-1]{
			res[index],res[index-1]=res[index-1],res[index       ]
			index--
		}
	}
}