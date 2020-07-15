package main

import "fmt"

type MaxQueue struct {
	Queue []int
	Help  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		Queue: []int{},
		Help:  []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Help) == 0 {
		return -1
	}
	return this.Help[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.Queue = append(this.Queue, value)

	for len(this.Help) != 0 && value > this.Help[len(this.Help)-1] {
		this.Help = this.Help[:len(this.Help)-1]
	}

	this.Help = append(this.Help, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Queue) == 0 {
		return -1
	}
	num := this.Queue[0]
	this.Queue = this.Queue[1:]
	if num == this.Help[0] {
		this.Help = this.Help[1:]
	}
	return num
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
func main() {
	obj := Constructor()
	fmt.Println(obj.Max_value())

	obj.Push_back(1)
	obj.Push_back(2)

	fmt.Println("max is ", obj.Max_value())
	fmt.Println("pop_front is ", obj.Pop_front())
	fmt.Println("max is ", obj.Max_value())
}
