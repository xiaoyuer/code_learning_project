/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || q== nil || q == nil {
        return root
    } 
     if q.Val < root.Val && p.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else if p.Val > root.Val && q.Val > root.Val{
        return lowestCommonAncestor(root.Right, p, q)
    } else {
        return root
    }
}