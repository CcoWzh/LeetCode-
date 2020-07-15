package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	maxBytes int                   //最大容量
	ll       *list.List            //双向链表
	cache    map[int]*list.Element //哈希表，指向链表元素
}

//链表中的值
type entry struct {
	key   int
	value int
}

//初始化
func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxBytes: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		//将元素e移动到list l的首部，如果e不属于list l，则list不改变。
		this.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if ele, ok := this.cache[key]; ok {
		//如果键存在，将该节点移到最前面，并更新对应节点的值
		this.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
	} else {
		if this.maxBytes == this.ll.Len() {
			//如果超过最大缓存容量，则删除最后的元素
			last := this.ll.Back()
			kv := last.Value.(*entry)
			this.ll.Remove(last)
			fmt.Println("删除的键值是：", kv)
			delete(this.cache, kv.key)
		}
		//不存在，则新增节点 ，在list的首部插入值为v的元素，并再哈希表中记录新增键
		ele := this.ll.PushFront(&entry{key, value})
		this.cache[key] = ele
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.printAll()
	lru.Put(2, 3)
	lru.printAll()
	lru.Put(4, 1)
	lru.printAll()
	lru.Put(3, 3)
	lru.printAll()
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

}

//打印链表中的所有值
func (this *LRUCache) printAll() {
	for e := this.ll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,无内容
	}
	fmt.Print("\n")
}
