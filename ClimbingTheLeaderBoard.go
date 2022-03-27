func climbingLeaderboard(ranked []int32, player []int32) []int {

    bantu := []int32{}
    bantu = append(bantu, ranked[0])
    // nilai := 0
    hasil := []int{}
    
    for _, j := range ranked{
        if j != bantu[len(bantu)-1]{
            bantu = append(bantu, j)
        }
    }
    
    urutan := len(player)-1
    for n, m := range bantu {

        if player[urutan] > m {
            hasil = append(hasil, n+1)
            fmt.Println("ranking :", n+1)
            fmt.Println("player: ", player[urutan])
            urutan = urutan - 1
        }
        
        if player[urutan] == m {
            hasil = append(hasil, n+1)
            fmt.Println("ranking :", n+1)
            fmt.Println("player: ", player[urutan])
            urutan = urutan - 1
        }
        
        if n == len(bantu) - 1 && player[urutan] < m{
            fmt.Println("ranking :", int32(n)+1)
            fmt.Println("player: ", player[urutan])
            hasil = append(hasil, n+2)
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(hasil)))

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
