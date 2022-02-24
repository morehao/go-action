package main

func main() {

}

func kthLargest(root *TreeNode, k int) (res int) {
	var recur func(root *TreeNode)
	recur = func(root *TreeNode) {
		if root == nil || k <= 0 {
			return
		}
		recur(root.Right)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		recur(root.Left)
	}
	recur(root)
	return
}
