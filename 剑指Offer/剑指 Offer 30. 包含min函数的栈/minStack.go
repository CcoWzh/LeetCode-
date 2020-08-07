package main

type MinStack struct {
	stack1 []int //主栈
	stack2 []int //辅助栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	m := len(this.stack2)
	//单调递增栈
	if m == 0 || this.stack2[m-1] >= x {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	n, m := len(this.stack1), len(this.stack2)
	if n <= 0 {
		return
	}
	//如果出栈元素是最小值的话，则辅助栈也要出栈
	if this.stack1[n-1] == this.stack2[m-1] {
		this.stack2 = this.stack2[:m-1]
	}
	this.stack1 = this.stack1[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack1)
	if n == 0 {
		return 0
	}
	return this.stack1[n-1]
}

func (this *MinStack) Min() int {
	m := len(this.stack2)
	if m == 0 {
		return 0
	}
	return this.stack2[m-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
func main() {
	minStack := Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-3)
	//minStack.Push(6)
	//fmt.Println(minStack.stack1)
	//fmt.Println(minStack.Min()) //--> 返回 -3.
	//minStack.Pop()
	//fmt.Println(minStack.Top()) //--> 返回 -3.
	//fmt.Println(minStack.Min()) //--> 返回 -3.
	//fmt.Println(minStack.stack1)
	minStack.Pop()
}
