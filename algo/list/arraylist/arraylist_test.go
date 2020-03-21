package arraylist

import (
	"testing"
)

func TestListNew(t *testing.T) {
	list1 := New()

	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New(1, "b")

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}

	if actualValue, ok := list2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

// 测试 add 方法
func TestList_Add(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if ac := list.Empty(); ac != false {
		t.Errorf("Got %v expected %v", ac, false)
	}

	if ac := list.Size(); ac != 3 {
		t.Errorf("Got %v expected %v", ac, 3)
	}
	if ac, ok := list.Get(2); ac != "c" || !ok {
		t.Errorf("Got %v expected %v", ac, "c")
	}
}
