package main

import "fmt"

type MinStack struct {
	Values []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Values = append(this.Values, val)
}

func (this *MinStack) Pop() {
	this.Values = this.Values[:len(this.Values)-1]
}

func (this *MinStack) Top() int {
	mini := this.Values[0]
	for _, v := range this.Values {
		if mini < v {
			mini = v
		}
	}
	return mini
}

func (this *MinStack) GetMin() int {
	mini := this.Values[0]
	for _, v := range this.Values {
		if mini > v {
			mini = v
		}
	}
	return mini
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())
}
