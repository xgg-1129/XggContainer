package Sort

import "reflect"

func InsertSort(arr []int,module Type)[]int{
	switch module {
	case Asc:
		return insertAscSort(arr)
	case Desc:
		return insertDescSort(arr)
	default:
		return nil
	}
}
//可以用二分查找改进
func insertAscSort(arr []int) []int {
	n:=len(arr)
	res:=make([]int,n)
	copy(res,arr)
	for i:=1;i<n;i++{
		index :=i
		//如果比前面的数小
		for index>0&& res[index] < res[index-1]{
			res[index],res[index-1]=res[index-1],res[index       ]
			index--
		}
	}
	return res
}

func insertDescSort(arr []int) []int {
	res := insertAscSort(arr)
	Reverse(res)
	return res
}
func Reverse(slice interface{}){
	s := reflect.Indirect(reflect.ValueOf(slice))
	i := 0
	j := s.Len() - 1
	for i < j {
		x, y := s.Index(i).Interface(),
			s.Index(j).Interface()
		s.Index(i).Set(reflect.ValueOf(y))
		s.Index(j).Set(reflect.ValueOf(x))
		i++
		j--
	}

}
