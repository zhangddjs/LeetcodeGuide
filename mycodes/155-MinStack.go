package main

type MinStack struct {
	Mins  []int
	Stack []int
}

func Constructor() MinStack {
	return MinStack{
		Mins:  make([]int, 0),
		Stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Mins) == 0 || this.Mins[len(this.Mins)-1] >= val {
		this.Mins = append(this.Mins, val)
	}
}

func (this *MinStack) Pop() {
	val := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if val == this.GetMin() {
		this.Mins = this.Mins[:len(this.Mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Mins[len(this.Mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
