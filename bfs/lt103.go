package bfs

/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

*/

func ZigtagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 层序遍历结果
	ret := make([][]int, 0)
	// 栈1存储当前层级节点
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	isOld := true
	for len(stack) > 0 {
		// 保存当前层级节点的值
		tmp := make([]int, 0)
		// 保存下一层级节点
		stack2 := make([]*TreeNode, 0)
		for len(stack) > 0 {
			// pop出栈顶
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 记录每层每个节点的val值
			tmp = append(tmp, node.Val)
			if isOld {
				// 先push左节点再push右节点
				if node.Left != nil { // push左节点到栈
					stack2 = append(stack2, node.Left)
				}
				if node.Right != nil { // push右节点到栈
					stack2 = append(stack2, node.Right)
				}
			} else {
				// 先push右节点再push左节点
				if node.Right != nil { // push右节点到栈
					stack2 = append(stack2, node.Right)
				}
				if node.Left != nil { // push左节点到栈
					stack2 = append(stack2, node.Left)
				}
			}
		}
		// 每层遍历完，转变一下状态区分奇数偶数层
		isOld = !isOld
		// swap栈1与栈2
		stack = stack2
		ret = append(ret, tmp)
	}
	return ret
}
