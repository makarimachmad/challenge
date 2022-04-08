func MissingInteger(A []int) int {
    // write your code in Go 1.4
    sort.Ints(A)
    minimal := 1

    for _, j := range A{
        if j == minimal {
             minimal = minimal + 1
        }
    }

    return minimal
}
