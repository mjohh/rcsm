package main
import "fmt"

func bubblesort(a []int){
    len := len(a)
    for i:=0; i<len-1; i++{
        for j:=i+1; j<len; j++{
            if a[i] > a[j]{
                a[i],a[j] = a[j],a[i]
            }
        }
    }
}

func main() {
    a := []int{2,3,4,1,0,2,6,2,3,4,8,0}
    fmt.Println(a)

    bubblesort(a)
    fmt.Println(a)

    bubblesort(a)
    fmt.Println(a)

    a = []int{1,2,9,7,6,5,4,3}

    bubblesort(a)
    fmt.Println(a)
}
