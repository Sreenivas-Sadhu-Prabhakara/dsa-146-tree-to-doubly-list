package main

import (
	"reflect"
	"testing"
)

func TestToDoublyList(t *testing.T) {
	root := tn(10, tn(6, tn(4, nil, nil), tn(8, nil, nil)), tn(14, tn(12, nil, nil), tn(16, nil, nil)))
	head := ToDoublyList(root)

	// Forward walk is the inorder sequence.
	fwd := []int{}
	var tail *TreeNode
	for n := head; n != nil; n = n.Right {
		fwd = append(fwd, n.Val)
		tail = n
	}
	if !reflect.DeepEqual(fwd, []int{4, 6, 8, 10, 12, 14, 16}) {
		t.Errorf("forward = %v", fwd)
	}

	// Backward walk via Left links must mirror it.
	bwd := []int{}
	for n := tail; n != nil; n = n.Left {
		bwd = append(bwd, n.Val)
	}
	if !reflect.DeepEqual(bwd, []int{16, 14, 12, 10, 8, 6, 4}) {
		t.Errorf("backward = %v", bwd)
	}

	if head.Left != nil || tail.Right != nil {
		t.Error("list ends must be nil-terminated")
	}

	if ToDoublyList(nil) != nil {
		t.Error("empty tree gives nil list")
	}
	single := ToDoublyList(tn(5, nil, nil))
	if single.Val != 5 || single.Left != nil || single.Right != nil {
		t.Error("single node list wrong")
	}
}
