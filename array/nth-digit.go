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


//optimised
func findNthDigit(n int) int {
    n -= 1
    for digit := 1; digit < 11; digit++ {
        firstNum := math.Pow(10, float64(digit-1))
        if n - 9*int(firstNum)*digit < 0 {
            res, _ := strconv.Atoi(string(strconv.Itoa(int(firstNum)+n/digit)[n%digit]))
            return res
        }
        n -= 9*int(firstNum)*digit
    }
    return 0
}