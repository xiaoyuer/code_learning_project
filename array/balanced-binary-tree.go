func isBalanced(root *TreeNode) bool {
    return root == nil || isBalanced(root.Left) && isBalanced(root.Right) && math.Abs(height(root.Left) - height(root.Right)) < 2
}

func height(root *TreeNode) float64 {
    if root == nil {
        return 0
    }
    return math.Max(height(root.Left), height(root.Right)) + 1
}