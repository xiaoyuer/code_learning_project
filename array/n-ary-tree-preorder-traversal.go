/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 var res []int

 func preorder(root *Node) []int {
	 res = []int{}
	 dfs(root)
	 return res
 }
 
 func dfs(root *Node) {
	 if root != nil {
		 res = append(res, root.Val)
		 for _, n := range root.Children {
			 dfs(n)
		 }
	 }
 }


 //递归
 func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*Node{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children)-1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
 }