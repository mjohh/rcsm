package main
import "fmt"

func fibonacci(n int) []int{
    if n < 1 {
        return nil
    }else if n == 1 {
        s := []int{1}
        return s
    } else if n == 2 {
        s := []int{1,1}
        return s
    } else{
        s := make([]int,n)
        s[0] = 1
        s[1] = 1
        for i:=2; i<n; i++{
            s[i] = s[i-1]+s[i-2]
        }
        return s
    }
    return nil
}

func main(){
    s := fibonacci(0)
    fmt.Println(s)
    s = fibonacci(1)
    fmt.Println(s)
    s = fibonacci(2)
    fmt.Println(s)
    s = fibonacci(6)
    fmt.Println(s)
    fmt.Println("\nsss")
}
