# 数组 链表 跳表

## container-with-most-water

两边往中间收：

* 终止条件：i和j相遇
* 小的一边柱子往里移动

```go
func maxArea(height []int) int {
    o := 0
    i, j := 0, len(height)-1
    for i != j {
       
        s := (j-i)*min(height[i], height[j])
        if s > o {
            o = s
        }
        if height[i] > height[j] {
            j--
        } else {
            i++
        }
    }
    return o
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```

## move-zeroes

序列从左到右轮询遍历一遍，遍历过程中，序列分为两部分：已遍历序列，未遍历序列。 已遍历序列又分成两部分：左边的非0序列，右边的0序列。一开始两个序列长度都是0。

整个遍历过程可看作，将遍历到的一个新节点，插入到非0序列与0序列之间。

这个过程，对0元素，不需要操作，直接就放在0序列队尾；对非0元素，将其与0序列头交换 所需要的变量，除了遍历索引，需要一个lNonZ记录当前非0序列的长度，它可以定位非0序列的尾部和0序列头部。

```go
func moveZeroes(nums []int)  {
    j := 0
    for i:= range(nums) {
        if (nums[i] != 0) {  //遇到非0数，将其添加到非0序列队尾，并更新长度
            if(i != j) { //当前数不在非0序列尾，说明含有0序列，将其与0序列头交换
                nums[j] = nums[i];
                nums[i] = 0;
            }
            j++;
        }
    }
}
```



