package main
import "fmt"

func min(s []int) int{
    min := s[0]
    for _,v := range s{
        if min > v{
            min = v
        }
    }
    return min
}

func max(s []int) int{
    max := s[0]
    for _,v := range s{
        if max < v{
            max = v
        }
    }
    return max
}

func main(){
    a := []int{3,1,5,6,34,23,2}
    mi := min(a)
    mx := max(a)
    fmt.Println(a, mi, mx)

}
