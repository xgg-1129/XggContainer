package HashTable

import "unsafe"

type hmap struct {
	count int
	//hamp的容量是2的B次方，这么设计原因是能确保每个桶都能被访问到
	B     int

	//buckets指向底层的哈希表
	buckets unsafe.Pointer
	//golang的map采取渐进式的扩容机制，需要一个oldbuckets指向旧的桶，从而完成渐进式扩容
	oldbuckets unsafe.Pointer
	flag uint8  //说明hamp的状态
	nevacuate uintptr//下一个要迁移的酒桶下标
	
}

type myMap *hmap