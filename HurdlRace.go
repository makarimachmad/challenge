func hurdleRace(k int32, height []int32) (abc int32) {
    // Write your code here
    tertinggi := 0
    for _, j := range height{
        if j > k && int32(tertinggi) < j{
            tertinggi = int(j)
        }
    }
    
    if tertinggi != 0{
        abc = int32(tertinggi) - k
    }else{
        abc = int32(tertinggi)
    }
    return abc
}
