package main
import "fmt"

type I interface {
    Get() int
    Put(int)
}

func g(a interface{}) int {
    if a, ok := a.(I); ok{
        return a.Get()
    }
    return 0
}

type Age int
func (a Age) Get() int{
    return 99
}

func (a Age) Put(i int){
    
}
// also ok
/* 
func (a Age) Get() int{
    return 99
}

func (a Age) Put(i int){
    
} 
*/

func main(){
    s := 0
    fmt.Println(g(s))

    var age Age
    fmt.Println(g(&age))
    // *T could use T & *T methods
    // T could only use T methods
    // for *T methods may change vals which need arg by reference(or pointer) 
}
