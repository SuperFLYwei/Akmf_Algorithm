package set

import "testing"

func TestNew(t *testing.T) {
	set := New(1, 2, 3)
	if set.Len() != 3 {
		t.Errorf("expecting 3 elements in the set but got %v: %v", set.Len(), set.GetItems())
	}
	for _, n := range []int{1, 2, 3} {
		if !set.In(n) {
			t.Errorf("expecting %d to be present in the set but was not; set is %v", n, set.GetItems())
		}
	}
	if set.In(5) {
		t.Errorf("expecting 5 not to be present in the set but it was; set is %v", set.GetItems())
	}
}