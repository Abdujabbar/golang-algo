package StackViaSlice

//StackVS stack realized via slice
type StackVS struct {
	Items []interface{}
}

//InitStackVS initilize stack via slice
func InitStackVS() *StackVS {
	a := StackVS{}
	return &a
}

//Push pushes element to stack
func (a *StackVS) Push(element interface{}) {
	a.Items = append(a.Items, element)
}

//Pop pops element from stack
func (a *StackVS) Pop() interface{} {
	if a.IsEmpty() {
		return "STack is empty"
	}
	v := a.Items[len(a.Items)-1]
	a.Items = a.Items[:len(a.Items)-1]
	return v
}

//Size return size of stack
func (a *StackVS) Size() int {
	return len(a.Items)
}

//IsEmpty checks is empty the stack
func (a *StackVS) IsEmpty() bool {
	return a.Size() == 0
}
