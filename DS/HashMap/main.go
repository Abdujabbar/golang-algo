package HashMap

import "fmt"

//HashMap Data Structure
type HashMap struct {
	Collection map[int]map[string]interface{}
	Size       int
}

//InitHashMap initilize hashmap ds
func InitHashMap(sz int) *HashMap {
	a := HashMap{}
	a.Collection = make(map[int]map[string]interface{})
	a.Size = sz
	return &a
}

func hash(v string, sz int) int {
	sm := 0
	for _, code := range v {
		sm += int(code)
	}
	return sm % sz
}

// Add method
func (a *HashMap) Add(key string, v interface{}) {

	index := hash(key, a.Size)
	if vv, ok := a.Collection[index]; !ok {
		a.Collection[index] = make(map[string]interface{})
		a.Collection[index][key] = v
	} else {
		vv[key] = v
	}
}

//Lookup finding item by key
func (a *HashMap) Lookup(key string) interface{} {
	index := hash(key, a.Size)
	if vv, ok := a.Collection[index]; ok {
		if v, ok := vv[key]; ok {
			return v
		}
	}
	return nil
}

//Remove from HashMap
func (a *HashMap) Remove(key string) bool {
	index := hash(key, a.Size)
	if vv, ok := a.Collection[index]; ok {
		if _, ok := vv[key]; ok {
			delete(vv, key)
			return true
		}
	}
	return false
}

//Print print hashmap
func (a *HashMap) Print() {
	fmt.Println(a.Collection)
}
