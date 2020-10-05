package tree

/*
	给定一个二叉树进行后序遍历（非递归）
*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	rst := make([]int, 0)
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		// 1）将左节点放入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		// 2）将右节点放入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 3）把left和right压入栈顺序交换，并按前序遍历，则是后序遍历结果，然后进行反转
	for i := len(ret) - 1; i >= 0; i-- {
		rst = append(rst, ret[i])
	}
	return rst
}
