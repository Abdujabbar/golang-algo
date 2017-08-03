package Stack

import "testing"

var st Stack

func TestSize(t *testing.T) {
	st.Push(10)
	st.Push(20)
	if st.Size() != 2 {
		t.Error("Error on counting")
	}
	st.Clear()

	if st.Size() != 0 {
		t.Error("Error on counting after Clear")
	}
}
