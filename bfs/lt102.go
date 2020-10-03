package bfs

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue1 := make([]*TreeNode, 0)
	for len(queue1) > 0 {
		tmp := make([]int, 0)
		queue2 := make([]*TreeNode, 0)
		for len(queue1) > 0 {
			// pop 出一个元素
			node := queue1[0]
			queue1 = queue1[1:]
			tmp = append(tmp, node.Val)
			// push进左节点
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			// push进右节点
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}
		ret = append(ret, tmp)
		// swap两个队列，当当前层遍历完之后
		queue1 = queue2
	}
	return ret
}
