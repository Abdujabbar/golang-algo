package Sets

import "fmt"

// Set DS
type Set struct {
	Items map[interface{}]interface{}
}

//InitSet method for initilize
func InitSet() *Set {
	a := Set{}
	a.Items = make(map[interface{}]interface{})
	return &a
}

//Add adds element to set
func (a *Set) Add(element interface{}) bool {
	if _, ok := a.Items[element]; ok {
		return false
	}
	a.Items[element] = element
	return true
}

//Remove removes item from set
func (a *Set) Remove(element interface{}) bool {
	if _, ok := a.Items[element]; ok {
		delete(a.Items, element)
		return true
	}
	return false
}

//Print elements
func (a *Set) Print() {
	fmt.Println(a.Items)
}

//Size size of set
func (a *Set) Size() int {
	return len(a.Items)
}

// Union of sets
func (a *Set) Union(b *Set) *Set {
	c := InitSet()
	for _, v := range a.Items {
		c.Add(v)
	}
	for _, v := range b.Items {
		c.Add(v)
	}
	return c
}

//Intersect of set
func (a *Set) Intersect(b *Set) *Set {
	c := InitSet()
	for _, v := range a.Items {
		if _, ok := b.Items[v]; ok {
			c.Add(v)
		}
	}
	return c
}

// Difference of sets
func (a *Set) Difference(b *Set) *Set {
	c := InitSet()
	for _, v := range a.Items {
		if _, ok := b.Items[v]; !ok {
			c.Add(v)
		}
	}

	for _, v := range b.Items {
		if _, ok := a.Items[v]; !ok {
			c.Add(v)
		}
	}
	return c
}

//Subset method for checking subsets
func (a *Set) Subset(b *Set) bool {
	for _, v := range b.Items {
		if _, ok := a.Items[v]; !ok {
			return false
		}
	}
	return true
}
