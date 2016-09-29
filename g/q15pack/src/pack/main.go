package main
import "pack/stack"
import "fmt"


func main(){
    s :=  stack.Newstack(20)

    s.Push(1)
    s.Push(3)
    s.Push(0)
    s.Push(22)
    s.Push(2)

    fmt.Println(s)
}
