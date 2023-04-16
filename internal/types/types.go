package types

type GraphNode[T comparable] struct {
	Value     T
	Neighbors []*GraphNode[T]
	Costs     map[*GraphNode[T]]uint
}

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type ListNode[T comparable] struct {
	Value T
	Prev  *ListNode[T]
	Next  *ListNode[T]
}
