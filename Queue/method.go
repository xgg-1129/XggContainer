package Queue

import "fmt"

func GetQueue() queue {
	return queue{arr: make([]Object, 0)}
}

func (this *queue) IsEmpty() bool {
	return len(this.arr) == 0
}

func (this *queue) EnQueue(data Object) {
	this.arr = append(this.arr, data)
}

func (this *queue) DeQueue() bool {
	if len(this.arr) == 0 {
		return false
	} else {
		this.arr = this.arr[1:]
		return true
	}
}

func (this *queue) PrintQueue() {
	n := len(this.arr)
	for i := 0; i < n; i++ {
		fmt.Printf("%v\t", this.arr[i])
	}
}
