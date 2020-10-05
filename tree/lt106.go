package tree

/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

func buildTreez(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	node := &TreeNode{Val: val}
	i := 0
	for k, v := range inorder {
		if v == val {
			i = k
			break
		}
	}
	node.Left = buildTreez(inorder[:i], postorder[:i])
	node.Right = buildTreez(inorder[i+1:], postorder[i:len(postorder)-1])
	return node
}
