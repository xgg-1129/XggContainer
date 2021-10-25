package Sort

import (
	"fmt"
	"testing"
)

func TestAscSort(t *testing.T) {
	arr := []int{5,3,6,1,1,9}
	sort := InsertSort(arr,Desc)
	fmt.Println(arr)
	fmt.Println(sort)
}
