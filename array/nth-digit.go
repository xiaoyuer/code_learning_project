func findNthDigit(n int) int {
    digit, flag := 1, 9

    for n - digit*flag > 0 {
        n -= digit*flag
        digit++
        flag *= 10
    }
    if digit == 1 {
        return n
    }
    number := 1
    for i := 1; i < digit; i++ {
        number *= 10
    }
    number = number + (n-1)/digit
    idx := (n-1)%digit
    str := strconv.Itoa(number)
    res, _ := strconv.Atoi(string(str[idx]))
    return res
}