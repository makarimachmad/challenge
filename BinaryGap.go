func BinaryGap(N int) int {
    // write your code in Go 1.4
    n := int64(N)
    binary := strconv.FormatInt(n, 2) // 1111011
    maks := 0
    last := binary[0]
    tampung := []int{}

    for i, j := range binary{
        if i != 0{
            if string(last) == "1"{
                if string(j) == "0"{
                    maks = maks + 1
                }
                if string(j) == "1"{
                    tampung = append(tampung, maks)
                    maks = 0
                }
            }
        }
        
    }

    if tampung == nil{
        return 0
    }

    nilai := 0

    for _, j := range tampung{
        if j>nilai{
            nilai = j
        }
    }

    return int(nilai)
}
