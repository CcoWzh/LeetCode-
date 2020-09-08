package main

import "fmt"

type MinStack struct {
	stack []int //主栈
	help  []int //辅助栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		help:  make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.help) == 0 || x < this.help[len(this.help)-1] {
		this.help = append(this.help, x)
	} else { //如果大于栈顶元素的话
		top := this.help[len(this.help)-1]
		this.help = append(this.help, top)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.help = this.help[:len(this.help)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.help[len(this.help)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
