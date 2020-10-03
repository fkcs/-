package bfs

/*
给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:

    2
   / \
  1   3

输出:
1
*/

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	left := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left = node.Val
		// 队列存放顺序，先放由节点再放左节点，那么最后的节点一定是最下层的最左节点
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return left
}
