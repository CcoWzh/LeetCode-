### `LRU`缓存机制

运用你所掌握的数据结构，设计和实现一个  [LRU (最近最少使用) 缓存机制](https://baike.baidu.com/item/LRU)。它应该支持以下操作： 获取数据 `get` 和 写入数据 `put` 。

获取数据 `get(key)` - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 `put(key, value)` - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

**进阶:**

你是否可以在 **O(1)** 时间复杂度内完成这两种操作？

**示例:**

```go
LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
```

思路：

1. 首先确定好数据结构，根据LRU的描述，我们使用`hashMap`和`双向链表`，实现`O(1)`的操作，此外，我们需要定义2个计数器，表示当前容量和可以使用的总容量;

2. 初始化，返回LRU对象

3. 在`get`操作时，我们查询`map`，如果在`map`中存在，则将查询的值提到最前面，并返回查询到的值，否则返回-1

4. 在`put`操作时，我们：

   4.1. 我们先查看put更新的数值是否存在缓存中，如果在，则直接更新数值，并且把这个节点提到最前面

   4.2. 如果不在缓存中，则先查看当前容量是否已经满了，如果满了的话，则删除最后一个节点，同时删除`map`里的值；接着，建立`map`，插入数值，并且提到最前面。

golang代码实现：

```go
type LRUCache struct {
    capacity int
    len int
    hashMap map[int]*Node
    head  *Node //这个应该是一个空节点
    tail  *Node
}

type Node struct{
    Prev *Node
    Next *Node
    Val int
    Key int
}


func Constructor(capacity int) LRUCache {
    m:=make(map[int]*Node)
    lru:= LRUCache{capacity:capacity,hashMap:m,head:&Node{},tail:&Node{}}
    lru.head.Next=lru.tail //双向循环
    lru.tail.Prev=lru.head
    return lru
}


func (this *LRUCache) Get(key int) int {
    if v,ok:=this.hashMap[key];ok{ //双向链表的好处就是，可以直接把节点删除，然后放到最前面
       v.Prev.Next=v.Next
       v.Next.Prev=v.Prev
       n:=this.head.Next
       this.head.Next=v
       v.Prev=this.head
       n.Prev=v
       v.Next=n
       return v.Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if v,ok:=this.hashMap[key];ok{
       v.Prev.Next=v.Next
       v.Next.Prev=v.Prev
       n:=this.head.Next
       this.head.Next=v
       v.Prev=this.head
       n.Prev=v
       v.Next=n 
       v.Val=value
       return  
    }
    if this.len<this.capacity{
       this.len++
       node:=&Node{
           Val:value,
           Key:key,
       }
       this.hashMap[key]=node
       n:=this.head.Next
       this.head.Next=node
       node.Prev=this.head
       node.Next=n
       n.Prev=node
    }else{
        t:=this.tail.Prev
        this.tail.Prev.Prev.Next=this.tail
        this.tail.Prev= this.tail.Prev.Prev
        t.Val=value
        delete(this.hashMap,t.Key)
        t.Key=key
        this.hashMap[key]=t
        hn:=this.head.Next
        this.head.Next=t
        t.Prev=this.head
        t.Next=hn
        hn.Prev=t
    }
    return
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```



参考：

- [Go 刷 LeetCode 系列：经典（1） LRU缓存机制](https://mp.weixin.qq.com/s/j-vZBzBU4tXlnWOzoJA8TQ)