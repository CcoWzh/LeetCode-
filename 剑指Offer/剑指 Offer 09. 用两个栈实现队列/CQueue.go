package main

import "fmt"

//CQueue 定义两个栈
type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

//AppendTail stack1负责添加数据
func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

//stack2负责将stack1中的数据正序出队
func (this *CQueue) DeleteHead() int {
	//如果stack2中没有数据，则将stack1中栈全部弹出，装入stack2中
	if len(this.stack2) != 0 {
		n := len(this.stack2)
		//stack2出栈
		temp := this.stack2[n-1]
		this.stack2 = this.stack2[:n-1]
		return temp
	} else {
		for len(this.stack1) != 0 {
			s := len(this.stack1)
			this.stack2 = append(this.stack2, this.stack1[s-1])
			//stack1出栈
			this.stack1 = this.stack1[:s-1]
		}
	}
	n := len(this.stack2)
	if n == 0 {
		return -1
	} else {
		temp := this.stack2[n-1]
		this.stack2 = this.stack2[:n-1]
		return temp
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func main() {
	obj := Constructor()
	param_1 := obj.DeleteHead()
	fmt.Println(param_1)

	obj.AppendTail(5)
	obj.AppendTail(2)
	param_2 := obj.DeleteHead()
	fmt.Println(param_2)

	param_3 := obj.DeleteHead()
	fmt.Println(param_3)
}
