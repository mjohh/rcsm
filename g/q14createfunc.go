package main
import "fmt"

func plusTwo() func (int) int{
    return func (x int) int{
        return x+2
    }
}


func plusX() func (a int, x int) int{
    return func (a int, x int) int{
        return a+x
    }
}

func main(){
    f := plusTwo()
    fmt.Println(3,f(3))

    f2 := plusX()
    fmt.Println(3,6,f2(3,6))
}
