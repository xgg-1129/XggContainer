package HashTable

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"sync"
)


const (
	ExpandFactor = 0.75
	ExpandSize=2
	InitCapacity = 23
)

type HashMap struct {
	array []*Node
	capacity int
	len int

	//考虑并发安全
	mu sync.RWMutex
}

func (m *HashMap) Capacity() int {
	return m.capacity
}

func (m *HashMap) Len() int {
	return m.len
}

func NewHashMap(cap int) *HashMap{
	h := &HashMap{}
	if cap<InitCapacity {
		h.capacity=InitCapacity
		h.array = make([]*Node, InitCapacity, InitCapacity)
	}else{
		h.capacity=cap
		h.array = make([]*Node, cap)
	}
	return h

}
//index= Hash(key) % size
func (m *HashMap) Index(key string) uint64{
	//求哈希
	hash:=m.hash([]byte(key))
	//求下标
	return hash% uint64(m.capacity)
}
func (m *HashMap) hash(key []byte)uint64{
	h := xxhash.New64()
	h.Write(key)
	return 	h.Sum64()
}

func (m *HashMap) Get(key string) interface{}{
	m.mu.RLock()
	defer m.mu.RUnlock()
	index:=m.Index(key)

	head:=m.array[index]
	for head!=nil{
		if head.key==key{
			return head.value
		}
		head=head.next
	}
	return nil
}

//头插法，速度最快
func (m *HashMap) Put(key string,value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	//如果超过负载因子，则对齐扩容
	newLen:=m.len+1
	if float64(m.len)/float64(m.capacity) > ExpandFactor{
		//不些cap，会默认len==cap
		NewHm:=NewHashMap(m.capacity*2)
		//遍历旧的Array
		for _,item := range m.array{
			for item!=nil {
				NewHm.Put(key,value)
				item=item.next
			}
		}
		m=NewHm
	}
	m.len=newLen
	index:=m.Index(key)
	temp:=m.array[index]
	m.array[index]=&Node{
		key:   key,
		value: value,
	}
	m.array[index].next=temp
}

func (m *HashMap) Delete(key string){
	if node := m.Get(key);node==nil{
		return
	}
	//不能放前面，不然会和Get冲突
	m.mu.Lock()
	defer m.mu.Unlock()
	index:=m.Index(key)
	head:=m.array[index]
	//如果头结点就是要删除的结点
	if head.key== key{
		m.array[index]=head.next
		m.len--
		return
	}
	for head.next!=nil{
		if head.next.key==key{
			head.next=head.next.next
			m.len--
			return
		}
		head=head.next
	}
}

func (m *HashMap) PrintMap() {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for index,item := range m.array{
		for item!=nil{
			fmt.Printf("[index=%d, key=%s ,value=%s] ",index,item.key,item.value)
			item=item.next
		}
		fmt.Println()
	}
}
