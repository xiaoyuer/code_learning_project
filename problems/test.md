# Algorithms

### Fibbinacci数列

### 链表有无环

### 树（前序，中序，后序遍历）

### 归并排序 冒泡排序

### 开根号





## 

### 合并两个有序列表 从小到大

```text
func merge(nums1 []int, m int, nums2 []int, n int)  {
	temp := make([]int, m)
	copy(temp, nums1)
	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
```

```text
Func sort(nums []int) []int {
n := len(nums)
for i := 0 ; i< n; i++ {
for j := i; j < n; j++ {
if nums[i] >nums[j]{
nums[i],nums[j] = nums[j],nums[i]
}
}
}
return nums
}
```

```text
//归并排序
func sort(nums []int) []int {
var res []int
flag := len(nums)/2
if flag > 0 {
nums1 := sort(nums[0:flag])
nums2 := sort(nums[flag:])
res = sortLess(nums1, nums2)
return res
}
return nums
}

func sortLess(left, right []int) []int {
var res []int
m,n := len(left), len(right)
for i,j:=0,0; i<m, j<n {
if left[i] > right[j] {
res = append(res, right[j])
j++
} else {
res = append(res, left[i])
i++
}
}
return res
}

```

二叉树中序遍历

```text
type treeNode struct {
Val int
Left *treeNode
Right *treeNode
}

func binaryTree(root *treeNode) []int {
       var res []int
for root.Left != nil && root.right != nil {
res = append(res, binaryTree(root.left)...)
res = append(res,root.Val)
res = append(res, binaryTree(root.right)...)
}
return res
}
```

```text
//二叉树叶子结点判断是否为等差数列（遍历子节点）（遍历可用递归和迭代）
func binaryTree(root *TreeNode)bool {
var res []int
for root.left != nil && root.Right != nil {
res = append(res, 
}
return compare(res)
}

func leaf(root *treeNode) []int {
var res []int
If root.Left == nil &&  root.right == nil {
res = append(root.Val)
}
return res
}

func compare(nums []int) bool {
If len(nums) <= 2 {
return true
}
for i,j,k := 0,1,2; k < len(nums); i++,j++,k++ {
if (nums[j]-nums[i]) != (nums[k]-nums[j]) {
return false
}
}
return true
}

```

当前表中字段为 a, b, c, d, ...

```text
SELECT * FROM t WHERE a=? AND b=? OR c=? ORDER BY d
```



## 实际问题

### 红包1元随机分给M个人



