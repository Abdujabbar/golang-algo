package QueueViaSlice

//QueueVS structure queue built in slice
type QueueVS struct {
	Items []interface{}
}

//InitQueueVS initilize QueueVS
func InitQueueVS() *QueueVS {
	a := QueueVS{}
	a.Items = make([]interface{}, 0)
	return &a
}

//Enqueue on queue
func (a *QueueVS) Enqueue(element interface{}) {
	a.Items = append(a.Items, element)
}

//Dequeue on queue
func (a *QueueVS) Dequeue() interface{} {
	if a.IsEmpty() {
		return "Queue is empty"
	}
	top := a.Items[0]
	a.Items = a.Items[1:]
	return top
}

//Size returns size of queue
func (a *QueueVS) Size() int {
	return len(a.Items)
}

//IsEmpty checking is empty
func (a *QueueVS) IsEmpty() bool {
	return a.Size() == 0
}

//Clear clearing queue
func (a *QueueVS) Clear() {
	a.Items = make([]interface{}, 0)
}
