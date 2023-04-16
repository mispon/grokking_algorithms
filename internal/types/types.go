package types

type GraphNode[T any] struct {
	Value     any
	Neighbors []*GraphNode[T]
}

type TreeNode[T any] struct {
	Value any
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type ListNode[T any] struct {
	Value any
	Prev  *ListNode[T]
	Next  *ListNode[T]
}
