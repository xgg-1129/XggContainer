package Sort


/*无法处理负数*/
func RadixSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return radixAscSort(arr)
	case Desc:
		return radixDescSort(arr)
	default:
		return nil
	}
}

func getDigits(num int)int{
	switch  {
	case num<10:
		return 1
	case num<100:
		return 2
	case num<1000:
		return 3
	case num<10000:
		return 4
	case num<100000:
		return 5
	case num<1000000:
		return 6
	default:
		return -1 //error
	}
}

func radixAscSort(arr []int) []int {
	n:=len(arr)
	res:=make([]int,n)
	if n==0||n==1{
		return  res
	}
	copy(res,arr)
	//求出最大的位数
	max := 0
	for _,item := range arr{
		digit:=getDigits(item)
		if max<getDigits(item){
			max=digit
		}
	}
	//创建buckets数组
	buckets:=make([][]int,10)
	//开始逐位比较
	for i,mod,dev:=0,10,1;i<max;i++{
		for _,value := range res{
			buckets[(value/dev)%mod]=append(buckets[(value/dev)%mod],value)
		}
		res =make([]int,0)
		for _,bucket :=range buckets{
			for _,v :=range bucket{
				res=append(res,v)
			}
		}
		buckets =make([][]int,10)
		mod*=10
		dev*=10
	}
	buckets=nil
	return res
}


func radixDescSort(arr []int) []int {
	res:=radixAscSort(arr)
	Reverse(res)
	return res
}
