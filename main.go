package main

import "fmt"

// TreeNode doubles as a doubly linked list node after conversion:
// Left becomes prev, Right becomes next.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// ToDoublyList rewires the tree, in place, into a sorted-order (inorder)
// doubly linked list and returns its head. An inorder walk threads each
// visited node to the previously visited one.
func ToDoublyList(root *TreeNode) *TreeNode {
	var head, prev *TreeNode
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		if prev == nil {
			head = n // smallest node: the list's head
		} else {
			prev.Right = n
			n.Left = prev
		}
		prev = n
		walk(n.Right)
	}
	walk(root)
	if head != nil {
		head.Left = nil
	}
	if prev != nil {
		prev.Right = nil
	}
	return head
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(10, tn(6, tn(4, nil, nil), tn(8, nil, nil)), tn(14, tn(12, nil, nil), tn(16, nil, nil)))
	for n := ToDoublyList(root); n != nil; n = n.Right {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}
