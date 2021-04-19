# Algorithms

### Fibbinacci数列

### 链表有无环

### 树（前序，中序，后序遍历）

### 归并排序 冒泡排序

### 开根号

  
p.p1 {margin: 0.0px 0.0px 0.0px 0.0px; font: 13.0px 'Helvetica Neue'}  
p.p2 {margin: 0.0px 0.0px 0.0px 0.0px; font: 13.0px 'Helvetica Neue'; min-height: 15.0px}  
p.p3 {margin: 0.0px 0.0px 0.0px 0.0px; font: 13.0px 'PingFang SC'}  
span.s1 {font: 13.0px 'Helvetica Neue'}  


array：

container with most water  11

3sum 15

climing stairs 70

Lru

move zeoros 283

linked list

1.reverse linked list 92  206  1265

2.swap node in pairs 24

链表环 141 142

25 reverse nodes 

栈

20 有效的括号

155 最小栈

84（hard）largest rectangel 

239 滑动窗口最大值

deque



map

242 有效的字母异位词

49 字母异位词

1 两数之和



tree

94 

144

590 n叉树

589

429 层



递归

70 爬楼梯

22 括号

反转二叉树 226

98 验证二叉搜索树

104 二叉树的最大深度

111 最小深度

297 二叉树序列化



分治

50 pow

78 子集

169多数元素（important）

17 电话号

51 n 皇后



深度优先和广度优先

102 二叉树层序遍历

22 括号生成

515 二叉树每一行最大值



贪心：860 柠檬水找零



二分查找

69  x的平方根

33 搜索旋转螺旋数组



动态规划

62 不同路径

63

1143 最长公共子序列

120 三角形最小路径和

53 最大子序和



trie树 

 208

212 单词搜索

并查集

200岛屿数量



avl树和红黑树

位运算

191 位1的个数

231 2的幂

190 颠倒二进制位

338 比特位计数

146 lru缓存



排序

1122 数组相对排序

242 有效的字母异位词

56 合并区间（important）

493 反转对



字符串

58 最后一个词的长度

387 第一个唯一字符

14 最长公共前缀

151 reverse word in string

557

917



125 验证回文串







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



