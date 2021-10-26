package Sort

import (
	"fmt"
	"testing"
)

func TestAscSort(t *testing.T) {
	arr := []int{9,3,6,1,1,44,19,0}
	sort := RadixSort(arr,Asc)
	fmt.Println(arr)
	fmt.Println(sort)

}
