func climbingLeaderboard(ranked []int, player []int32) []int {

    hasil := []int{}
    sort.Sort(sort.Reverse(sort.IntSlice(ranked)))
    
    bantu := []int{}
    bantu = append(bantu, ranked[0])
    
    
    for _, j := range ranked{
        if j != bantu[len(bantu)-1]{
            bantu = append(bantu, j)
        }
    }
    
    l := len(bantu)
    
    for _, s := range player{
        for l > 0 && int(s) >= bantu[l-1]{
            l = l-1
        }
        hasil = append(hasil, l+1)
    }

    return hasil

}
// input
// 7
// 100 100 50 40 40 20 10
// 4
// 5 25 50 120
//output
//6 4 2 1

//input
// 6
// 100 90 90 80 75 60
// 5
// 50 65 77 90 102
//output
// 6 5 4 2 1
