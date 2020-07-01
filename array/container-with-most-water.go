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