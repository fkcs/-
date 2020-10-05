package tree

// 二叉树中序遍历（非递归）
// 左中右
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	node := root
	rst := make([]int, 0)
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left // 左
		} else {
			// 出栈
			node = stack[len(stack)-1] // 中
			stack = stack[:len(stack)-1]
			rst = append(rst, node.Val)
			node = node.Right // 右
		}
	}
	return rst
}
