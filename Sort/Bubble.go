package Sort



type Type int

const (

	Asc Type = 0
	Desc Type =1
)
/*除了名字好听外一无是处的的排序算法*/

func BubbleSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return bubbleAscSort(arr)
	case Desc:
		return bubbleDescSort(arr)
	default:
		return nil
	}
}
//可以加个flag，检查每次遍历是否有产生交换，如果没有产生交换序列已经有序，无需遍历可直接return
func bubbleAscSort(arr []int)[]int{
	n:=len(arr)
	res:=make([]int,n)
	copy(res,arr)
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			//注意这里是> ，不是>=
			if res[j]>res[j+1]{
				res[j],res[j+1]=res[j+1],res[j]
			}
		}
	}
	return res
}

func bubbleDescSort(arr []int)[]int{
	n:=len(arr)
	res:=make([]int,n)
	copy(res,arr)
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if res[j]<res[j+1]{
				res[j],res[j+1]=res[j+1],res[j]
			}
		}
	}
	return res
}
