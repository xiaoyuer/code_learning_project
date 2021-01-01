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

